function getUntilLevel(devs, ids, depth) {
  let _devs = [];

  while (depth > 0) {
    let friends = [];

    for (let i = 0; i < devs.length; i++) {
      const dev = devs[i];

      if (ids.includes(dev.id)) {
        friends = friends.concat(dev.friends);
      }
    }

    depth--;

    if (depth > 0) {
      ids = friends;
    } else {
      _devs = friends;
    }
  }

  return _devs;
}

function unique(arr) {
  return Array.from(new Set(arr));
}

// Using previous values to improve performance
function getSocialBubbles(devs, id) {
  let f1 = unique(getUntilLevel(devs, [id], 1));
  let f2 = unique(getUntilLevel(devs, f1, 1));
  let f3 = unique(getUntilLevel(devs, f2, 1));

  return [f1, f2, f3];
}

// Using generic solution for any depth
function getSocialBubbles02(devs, id) {
  return [
    unique(getUntilLevel(devs, [id], 1)),
    unique(getUntilLevel(devs, [id], 2)),
    unique(getUntilLevel(devs, [id], 3))
  ];
}

const people = [
  { id: "0", friends: ["4", "1"] },
  { id: "1", friends: ["2", "3"] },
  { id: "2", friends: ["1"] },
  { id: "3", friends: ["2", "4"] },
  { id: "4", friends: ["3"] }
];

console.log(">> Social bubble 01 (Jão):", getSocialBubbles(people, "0"));
console.log(">> Social bubble 02 (Jão):", getSocialBubbles02(people, "0"));
