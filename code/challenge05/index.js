function BinarySearchTree(value) {
  this.value = value;
  this.left = null;
  this.right = null;
}

function fromArrayToBinarySearchTree(array) {
  array.sort()

  function _fromArrayToBinarySearchTree(array) {
    if (array.length === 0)
      return null

    let mid = Math.floor(array.length / 2)

    let root = new BinarySearchTree(array[mid]);

    root.left = _fromArrayToBinarySearchTree(array.slice(0, mid));
    root.right = _fromArrayToBinarySearchTree(array.slice(mid + 1));

    return root;
  }

  return _fromArrayToBinarySearchTree(array);
}

console.log(fromArrayToBinarySearchTree([2, 3, 1, 4, 6, 7]))
