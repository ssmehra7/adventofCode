import { input } from "./input.js";

function parseInput(input) {
  let arr = [];
  arr = input.split(`\n`);
  let newArr = arr.map((x) => {
    let y = x.split(" ").map(Number);
    return y;
  });

  return newArr;
}

function main() {
  let inputData = parseInput(input);

  let sum = 0;
  for (let i = 0; i < inputData.length; i++) {
    let increasing = true;
    let decreasing = true;

    let temp = inputData[i];

    for (let x = 1; x < temp.length; x++) {
      if (temp[x] < temp[x - 1]) {
        increasing = false;
      } else {
        decreasing = false;
      }
    }

    if (increasing || decreasing) {
      let isSafe = true;

      // Check the difference between adjacent levels
      for (let j = 1; j < temp.length; j++) {
        let diff = Math.abs(temp[j] - temp[j - 1]);
        if (diff < 1 || diff > 3) {
          isSafe = false;
          break;
        }
      }

      if (isSafe) {
        sum += 1;
      }
    }
  }

  console.log(sum);
}

main();
