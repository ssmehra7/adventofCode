import { input } from "./input.js";

function main() {
  const inputString = input;
  const inputArr = inputString.split(`\n`);

  let leftArr = [];
  let rightArr = [];

  for (let i = 0; i < inputArr.length; i++) {
    let infoString = inputArr[i];
    const tempArr = infoString.split("   ");

    leftArr.push(Number(tempArr[0]));
    rightArr.push(Number(tempArr[1]));
  }

  leftArr.sort((a, b) => a - b);
  rightArr.sort((a, b) => a - b);

  //   console.log(leftArr.length);
  //   console.log(rightArr.length);
  let sum = 0;

  for (let i = 0; i < leftArr.length; i++) {
    sum += Math.abs(leftArr[i] - rightArr[i]);
  }

  console.log(sum);
}

main();
