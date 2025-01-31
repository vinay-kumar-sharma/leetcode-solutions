// 827. Making A Large Island

const WATER = 0;
const LAND = 1;
const DIRS = [-1, 0, 0, -1, 0, 1, 1, 0];

/**
 * @param {number[][]} grid
 * @return {number}
 */
function largestIsland(grid) {
  const Y = grid.length;
  const X = grid[0].length;

  // Find islands
  let maxSize = 0;
  const islandSizes = [0, 0];
  for (let y = 0; y < Y; ++y) {
    for (let x = 0; x < X; ++x) {
      if (grid[y][x] == LAND) {
        const size = fill(grid, y, x, islandSizes.length);
        maxSize = Math.max(maxSize, size);
        islandSizes.push(size);
      }
    }
  }

  // If less than 2 islands found
  if (islandSizes.length < 4) {
    return maxSize + +(maxSize < Y * X);
  }

  // Check if islands can be connected
  const seen = new Array(islandSizes.length).fill(0);
  for (let y = 0; y < Y; ++y) {
    for (let x = 0; x < X; ++x) {
      if (grid[y][x] != WATER) {
        continue;
      }
      ++seen[0];
      let size = 1;
      for (let i = 0; i < DIRS.length; i += 2) {
        const y2 = y + DIRS[i];
        const x2 = x + DIRS[i + 1];
        if (y2 >= 0 && y2 < Y &&
          x2 >= 0 && x2 < X &&
          seen[grid[y2][x2]] != seen[0]
        ) {
          seen[grid[y2][x2]] = seen[0];
          size += islandSizes[grid[y2][x2]];
        }
      }
      maxSize = Math.max(maxSize, size);
    }
  }

  return maxSize;
};

/**
 * @param {number[][]} grid
 * @param {number} y
 * @param {number} x
 * @param {number} newValue
 * @return {number}
 */
function fill(grid, y, x, newValue) {
  const Y = grid.length;
  const X = grid[0].length;
  const stack = [y, x];
  const curValue = grid[y][x];
  grid[y][x] = newValue;

  let size = 0;
  do {
    ++size;
    x = stack.pop();
    y = stack.pop();
    for (let i = 0; i < DIRS.length; i += 2) {
      const y2 = y + DIRS[i];
      const x2 = x + DIRS[i + 1];
      if (y2 >= 0 && y2 < Y &&
        x2 >= 0 && x2 < X &&
        grid[y2][x2] == curValue
      ) {
        grid[y2][x2] = newValue;
        stack.push(y2, x2);
      }
    }
  } while (stack.length > 0);

  return size;
}