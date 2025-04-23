package Test

import (
	DS "DSA/DataStructures"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBtree(t *testing.T) {
	tree := DS.CreateBtree[int]()

	// Insert root
	root, err := tree.Root(50)
	assert.Nil(t, err)
	assert.Equal(t, 50, root.Value)

	// Insert nodes
	assert.Nil(t, tree.Insert(root, 30))
	assert.Nil(t, tree.Insert(root, 70))
	assert.Nil(t, tree.Insert(root, 20))
	assert.Nil(t, tree.Insert(root, 40))
	assert.Nil(t, tree.Insert(root, 60))
	assert.Nil(t, tree.Insert(root, 80))

	// Search tests
	node, err := tree.Search(root, 60)
	assert.Nil(t, err)
	assert.Equal(t, 60, node.Value)

	_, err = tree.Search(root, 100)
	assert.NotNil(t, err)

	// Delete leaf node (0 children)
	assert.Nil(t, tree.Delete(root, 20))
	_, err = tree.Search(root, 20)
	assert.NotNil(t, err)

	// Delete node with 1 child
	assert.Nil(t, tree.Delete(root, 30)) // 30 had only 40 after 20 was removed
	_, err = tree.Search(root, 30)
	assert.NotNil(t, err)

	// Delete node with 2 children
	assert.Nil(t, tree.Delete(root, 70)) // 70 has 60 and 80
	_, err = tree.Search(root, 70)
	assert.NotNil(t, err)

	// Final display
	tree.Display()
}

//! go test . /Tests/binarytree_test.go -v
