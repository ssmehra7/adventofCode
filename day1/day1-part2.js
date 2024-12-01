import { input } from "./input.js";

// const input = `3   4
// 4   3
// 2   5
// 1   3
// 3   9
// 3   3`;

function main() {
  const inputString = input;
  const inputArr = inputString.split(`\n`);

  let leftArr = [];
  let right = {};

  for (let i = 0; i < inputArr.length; i++) {
    let infoString = inputArr[i];
    const tempArr = infoString.split("   ");

    leftArr.push(Number(tempArr[0]));
    if (right[tempArr[1]] === undefined) {
      right[tempArr[1]] = 0;
    }
    right[tempArr[1]]++;
  }

  let sum = 0;

  for (let i = 0; i < leftArr.length; i++) {
    if (right[leftArr[i]] === undefined) {
      sum += 0;
    } else {
      sum += right[leftArr[i]] * leftArr[i];
    }
  }

  console.log(sum);
}

main();
