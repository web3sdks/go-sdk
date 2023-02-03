package merkle

import (
	"bytes"
	"crypto/rand"
	"errors"
	"math"
	"reflect"
	"sort"
	"sync"
)

const (
	// ModeProofGen is the proof generation configuration mode.
	ModeProofGen ModeType = iota
	// ModeTreeBuild is the tree building configuration mode.
	ModeTreeBuild
	// ModeProofGenAndTreeBuild is the proof generation and tree building configuration mode.
	ModeProofGenAndTreeBuild
	// Default hash result length using SHA256.
	defaultHashLen = 32
)

// ModeType is the type in the Merkle Tree configuration indicating what operations are performed.
type ModeType int

// DataBlock is the interface of input data blocks to generate the Merkle Tree.
type DataBlock interface {
	Serialize() ([]byte, error)
}

// HashFuncType is the signature of the hash functions used for Merkle Tree generation.
type HashFuncType func([]byte) ([]byte, error)

// Config is the configuration of Merkle Tree.
type Config struct {
	// Customizable hash function used for tree generation.
	HashFunc HashFuncType
	// Number of goroutines run in parallel.
	// If RunInParallel is true and NumRoutine is set to 0, use number of CPU as the number of goroutines.
	NumRoutines int
	// Mode of the Merkle Tree generation.
	Mode ModeType
	// If true, generate a dummy node with random hash value.
	// Otherwise, then the odd node situation is handled by duplicating the previous node.
	NoDuplicates bool
	// IMPORTANT: To match the behavior of merkletreejs, sort leaves before building tree
	SortLeaves bool
	// IMPORTANT: To match the behavior of merkletreejs, sort pairs before hashing them
	SortPairs bool
}

// MerkleTree implements the Merkle Tree structure
type MerkleTree struct {
	*Config
	// leafMap is the map of the leaf hash to the index in the Tree slice,
	// only available when config mode is ModeTreeBuild or ModeProofGenAndTreeBuild
	leafMap sync.Map
	// tree is the Merkle Tree structure, only available when config mode is ModeTreeBuild or ModeProofGenAndTreeBuild
	tree [][][]byte
	// Root is the Merkle root hash
	Root []byte
	// Leaves are Merkle Tree leaves, i.e. the hashes of the data blocks for tree generation
	Leaves [][]byte
	// Proofs are proofs to the data blocks generated during the tree building process
	Proofs []*Proof
	// Depth is the Merkle Tree depth
	Depth uint32
}

// Proof implements the Merkle Tree proof.
type Proof struct {
	Siblings [][]byte // sibling nodes to the Merkle Tree path of the data block
	Path     uint32   // path variable indicating whether the neighbor is on the left or right
}

// argType is used as the arguments for the handler functions when performing parallel computations.
// All the handler functions use this universal argument struct to eliminate interface conversion overhead.
// Each field in the struct may be used for different purpose in different handler functions,
// please refer to the comments at each handler function for details.
type argType struct {
	mt             *MerkleTree
	byteField1     [][]byte
	byteField2     [][]byte
	dataBlockField []DataBlock
	intField1      int
	intField2      int
	intField3      int
	intField4      int
	intField5      int
	uint32Field    uint32
}

// New generates a new Merkle Tree with specified configuration.
func New(config *Config, blocks []DataBlock) (m *MerkleTree, err error) {
	if len(blocks) <= 1 {
		return nil, errors.New("the number of data blocks must be greater than 1")
	}
	if config == nil {
		config = new(Config)
	}
	if config.HashFunc == nil {
		config.HashFunc = defaultHashFunc
	}
	// If the configuration mode is not set, then set it to ModeProofGen by default.
	if config.Mode == 0 {
		config.Mode = ModeProofGen
	}
	// If RunInParallel is true and NumRoutines is unset, then set NumRoutines to the number of CPU.

	m = &MerkleTree{Config: config}
	m.Depth = calTreeDepth(len(blocks))
	m.Leaves, err = m.leafGen(blocks)
	if err != nil {
		return
	}
	if m.Mode == ModeProofGen {
		err = m.proofGen()
		return
	}
	if m.Mode == ModeTreeBuild {
		err = m.treeBuild()
		return
	}
	if m.Mode == ModeProofGenAndTreeBuild {
		err = m.treeBuild()
		if err != nil {
			return
		}
		m.initProofs()
		for i := 0; i < len(m.tree); i++ {
			m.updateProofs(m.tree[i], len(m.tree[i]), i)
		}
		return
	}
	return nil, errors.New("invalid configuration mode")
}

// calTreeDepth calculates the tree depth,
// the tree depth is then used to declare the capacity of the proof slices.
func calTreeDepth(blockLen int) uint32 {
	log2BlockLen := math.Log2(float64(blockLen))
	// check if log2BlockLen is an integer
	if log2BlockLen != math.Trunc(log2BlockLen) {
		return uint32(log2BlockLen) + 1
	}
	return uint32(log2BlockLen)
}

func (m *MerkleTree) initProofs() {
	numLeaves := len(m.Leaves)
	m.Proofs = make([]*Proof, numLeaves)
	for i := 0; i < numLeaves; i++ {
		m.Proofs[i] = new(Proof)
		m.Proofs[i].Siblings = make([][]byte, 0, m.Depth)
	}
}

func (m *MerkleTree) proofGen() (err error) {
	numLeaves := len(m.Leaves)
	m.initProofs()
	buf := make([][]byte, numLeaves)
	copy(buf, m.Leaves)
	var prevLen int
	buf, prevLen, err = m.fixOdd(buf, numLeaves)
	if err != nil {
		return
	}
	m.updateProofs(buf, numLeaves, 0)
	for step := 1; step < int(m.Depth); step++ {
		for idx := 0; idx < prevLen; idx += 2 {
			buf[idx>>1], err = m.HashFunc(append(buf[idx], buf[idx+1]...))
			if err != nil {
				return
			}
		}
		prevLen >>= 1
		buf, prevLen, err = m.fixOdd(buf, prevLen)
		if err != nil {
			return
		}
		m.updateProofs(buf, prevLen, step)
	}
	m.Root, err = m.HashFunc(append(buf[0], buf[1]...))
	return
}

// proofGenHandler generates the proofs in parallel.
// arg fields:
//
//	mt: the Merkle Tree instance
//	byteField1: buf1
//	byteField2: buf2
//	intField1: start
//	intField2: prevLen
//	intField3: numRoutines
//
// return:
//
//	error
func proofGenHandler(arg argType) error {
	var (
		hashFunc    = arg.mt.HashFunc
		buf1        = arg.byteField1
		buf2        = arg.byteField2
		start       = arg.intField1
		prevLen     = arg.intField2
		numRoutines = arg.intField3
	)
	for i := start; i < prevLen; i += numRoutines << 1 {
		newHash, err := hashFunc(append(buf1[i], buf1[i+1]...))
		if err != nil {
			return err
		}
		buf2[i>>1] = newHash
	}
	return nil
}

// if the length of the buffer calculating the Merkle Tree is odd, then append a node to the buffer
// if AllowDuplicates is true, append a node by duplicating the previous node
// otherwise, append a node by random
func (m *MerkleTree) fixOdd(buf [][]byte, prevLen int) ([][]byte, int, error) {
	if prevLen&1 == 0 {
		return buf, prevLen, nil
	}
	var appendNode []byte
	if m.NoDuplicates {
		var err error
		appendNode, err = getDummyHash()
		if err != nil {
			return nil, 0, err
		}
	} else {
		appendNode = buf[prevLen-1]
	}
	prevLen++
	if len(buf) < prevLen {
		buf = append(buf, appendNode)
	} else {
		buf[prevLen-1] = appendNode
	}
	return buf, prevLen, nil
}

func (m *MerkleTree) updateProofs(buf [][]byte, bufLen, step int) {
	batch := 1 << step
	for i := 0; i < bufLen; i += 2 {
		m.updatePairProof(buf, i, batch, step)
	}
}

// updateProofHandler updates the proofs in parallel.
// arg fields:
//
//	mt: the Merkle Tree instance
//	byteField1: buf
//	intField1: start
//	intField2: batch
//	intField3: step
//	intField4: bufLen
//	intField5: numRoutines
//
// return:
//
//	nothing (nil)
func updateProofHandler(arg argType) error {
	var (
		mt          = arg.mt
		buf         = arg.byteField1
		start       = arg.intField1
		batch       = arg.intField2
		step        = arg.intField3
		bufLen      = arg.intField4
		numRoutines = arg.intField5
	)
	for i := start; i < bufLen; i += numRoutines << 1 {
		mt.updatePairProof(buf, i, batch, step)
	}
	// return the nil error to be compatible with the handler type
	return nil
}

func (m *MerkleTree) updatePairProof(buf [][]byte, idx, batch, step int) {
	start := idx * batch
	end := min(start+batch, len(m.Proofs))
	for i := start; i < end; i++ {
		m.Proofs[i].Path += 1 << step
		m.Proofs[i].Siblings = append(m.Proofs[i].Siblings, buf[idx+1])
	}
	start += batch
	end = min(start+batch, len(m.Proofs))
	for i := start; i < end; i++ {
		m.Proofs[i].Siblings = append(m.Proofs[i].Siblings, buf[idx])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// generate a dummy hash to make odd-length buffer even
func getDummyHash() ([]byte, error) {
	dummyBytes := make([]byte, defaultHashLen)
	_, err := rand.Read(dummyBytes)
	if err != nil {
		return nil, err
	}
	return dummyBytes, nil
}

// IMPORTANT: Do not call hash function on leaves!
func (m *MerkleTree) leafGen(blocks []DataBlock) ([][]byte, error) {
	var (
		lenLeaves = len(blocks)
		leaves    = make([][]byte, lenLeaves)
	)
	for i := 0; i < lenLeaves; i++ {
		data, err := blocks[i].Serialize()
		if err != nil {
			return nil, err
		}

		leaves[i] = data
	}

	// IMPORTANT: To match merkletreejs, we need to sort the leaves before hashing
	if m.SortLeaves {
		sort.Slice(leaves, func(i, j int) bool {
			return bytes.Compare(leaves[i], leaves[j]) < 0
		})
	}

	return leaves, nil
}

// leafGenHandler generates the leaves in parallel.
// arg fields:
//
//	mt: the Merkle Tree instance
//	byteField1: leaves
//	dataBlockField: blocks
//	intField1: start
//	intField2: lenLeaves
//	intField3: numRoutines
//
// return:
//
//	error
func leafGenHandler(arg argType) error {
	var (
		hashFunc    = arg.mt.HashFunc
		blocks      = arg.dataBlockField
		leaves      = arg.byteField1
		start       = arg.intField1
		lenLeaves   = arg.intField2
		numRoutines = arg.intField3
	)
	for i := start; i < lenLeaves; i += numRoutines {
		data, err := blocks[i].Serialize()
		if err != nil {
			return err
		}
		var hash []byte
		if hash, err = hashFunc(data); err != nil {
			return err
		}
		leaves[i] = hash
	}
	return nil
}

func (m *MerkleTree) treeBuild() (err error) {
	numLeaves := len(m.Leaves)
	finishMap := make(chan struct{})
	go func() {
		for i := 0; i < numLeaves; i++ {
			m.leafMap.Store(string(m.Leaves[i]), i)
		}
		finishMap <- struct{}{}
	}()
	m.tree = make([][][]byte, m.Depth)
	m.tree[0] = make([][]byte, numLeaves)
	copy(m.tree[0], m.Leaves)
	var prevLen int
	m.tree[0], prevLen, err = m.fixOdd(m.tree[0], numLeaves)
	if err != nil {
		return
	}
	for i := uint32(0); i < m.Depth-1; i++ {
		m.tree[i+1] = make([][]byte, prevLen>>1)

		for j := 0; j < prevLen; j += 2 {
			// VERY IMPORTANT: If the two leaves are duplicates (aka. an odd leaf in this case), we propagate it up
			// Instead of hashing the duplicates together. This is to match merkletreejs implementation.
			if (reflect.DeepEqual(m.tree[i][j], m.tree[i][j+1])) {
				m.tree[i+1][j>>1] = m.tree[i][j]
			} else {
				// IMPORTANT: To match merkletreejs, we sort the two leaves before hashing them together
				if m.SortPairs {
					if bytes.Compare(m.tree[i][j], m.tree[i][j+1]) < 0 {
						m.tree[i+1][j>>1], err = m.HashFunc(append(m.tree[i][j], m.tree[i][j+1]...))
					} else {
						m.tree[i+1][j>>1], err = m.HashFunc(append(m.tree[i][j+1], m.tree[i][j]...))
					}
				} else {
					m.tree[i+1][j>>1], err = m.HashFunc(append(m.tree[i][j], m.tree[i][j+1]...))
				}
			}

			if err != nil {
				return
			}
		}
		
		m.tree[i+1], prevLen, err = m.fixOdd(m.tree[i+1], len(m.tree[i+1]))
		if err != nil {
			return
		}
	}
	m.Root, err = m.HashFunc(append(m.tree[m.Depth-1][0], m.tree[m.Depth-1][1]...))
	if err != nil {
		return
	}
	<-finishMap
	return
}

// treeBuildHandler builds the tree in parallel.
// arg fields:
//
//	mt: the Merkle Tree instance
//	intField1: start
//	intField2: prevLen
//	intField3: numRoutines
//	uint32Field: depth
//
// return:
//
//	error
func treeBuildHandler(arg argType) error {
	var (
		mt          = arg.mt
		start       = arg.intField1
		prevLen     = arg.intField2
		numRoutines = arg.intField3
		depth       = arg.uint32Field
	)
	for i := start; i < prevLen; i += numRoutines << 1 {
		newHash, err := mt.HashFunc(append(mt.tree[depth][i], mt.tree[depth][i+1]...))
		if err != nil {
			return err
		}
		mt.tree[depth+1][i>>1] = newHash
	}
	return nil
}

// Verify verifies the data block with the Merkle Tree proof
func (m *MerkleTree) Verify(dataBlock DataBlock, proof *Proof) (bool, error) {
	return Verify(dataBlock, proof, m.Root, m.HashFunc)
}

// Verify verifies the data block with the Merkle Tree proof and Merkle root hash
func Verify(dataBlock DataBlock, proof *Proof, root []byte, hashFunc HashFuncType) (bool, error) {
	if dataBlock == nil {
		return false, errors.New("data block is nil")
	}
	if proof == nil {
		return false, errors.New("proof is nil")
	}
	if hashFunc == nil {
		hashFunc = defaultHashFunc
	}
	var (
		data, err = dataBlock.Serialize()
		hash      []byte
	)
	if err != nil {
		return false, err
	}
	hash, err = hashFunc(data)
	if err != nil {
		return false, err
	}
	path := proof.Path
	for _, n := range proof.Siblings {
		if path&1 == 1 {
			hash, err = hashFunc(append(hash, n...))
		} else {
			hash, err = hashFunc(append(n, hash...))
		}
		if err != nil {
			return false, err
		}
		path >>= 1
	}
	return bytes.Equal(hash, root), nil
}

// GenerateProof generates the Merkle proof for a data block with the Merkle Tree structure generated beforehand.
// The method is only available when the configuration mode is ModeTreeBuild or ModeProofGenAndTreeBuild.
// In ModeProofGen, proofs for all the data blocks are already generated, and the Merkle Tree structure is not cached.
func (m *MerkleTree) GenerateProof(dataBlock DataBlock) (*Proof, error) {
	if m.Mode != ModeTreeBuild && m.Mode != ModeProofGenAndTreeBuild {
		return nil, errors.New("merkle Tree is not in built, could not generate proof by this method")
	}
	blockByte, err := dataBlock.Serialize()
	if err != nil {
		return nil, err
	}
	var blockHash []byte
	if blockHash, err = m.HashFunc(blockByte); err != nil {
		return nil, err
	}
	val, ok := m.leafMap.Load(string(blockHash))
	if !ok {
		return nil, errors.New("data block is not a member of the Merkle Tree")
	}
	var (
		idx      = val.(int)
		path     uint32
		siblings = make([][]byte, m.Depth)
	)
	for i := uint32(0); i < m.Depth; i++ {
		if idx&1 == 1 {
			siblings[i] = m.tree[i][idx-1]
		} else {
			path += 1 << i
			siblings[i] = m.tree[i][idx+1]
		}
		idx >>= 1
	}
	return &Proof{
		Path:     path,
		Siblings: siblings,
	}, nil
}