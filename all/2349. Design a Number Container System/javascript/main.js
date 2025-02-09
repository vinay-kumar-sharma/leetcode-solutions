// 2349. Design a Number Container System

class NumberContainers {
	constructor() {
	  this.num_map = {};
	  this.num_map_cache = {};
	  this.idx_map = {};
	}
  
	/**
	* @param {number} index
	* @param {number} number
	* @return {void}
	*/
	change(index, number) {
	  if (index in this.idx_map) {
		const old_num = this.idx_map[index];
		if (old_num == number) return;
		const set = this.num_map[old_num];
		set.delete(index);
		if (old_num in this.num_map_cache &&
		  index == this.num_map_cache[old_num]) {
		  delete this.num_map_cache[old_num];
		}
	  }
	  this.idx_map[index] = number;
	  this.num_map[number] ||= new Set();
	  this.num_map[number].add(index);
	  if (number in this.num_map_cache) {
		this.num_map_cache[number] = Math.min(index, this.num_map_cache[number]);
	  }
	}
  
	/**
	* @param {number} number
	* @return {number}
	*/
	find(number) {
	  if (!this.num_map[number] || this.num_map[number].size <= 0) {
		return -1;
	  }
	  if (number in this.num_map_cache) {
		return this.num_map_cache[number];
	  }
	  const arr = [...this.num_map[number].values()].sort((a, b) => a - b);
	  this.num_map_cache[number] = arr[0];
	  return arr[0];
	}
  }
  
  /** 
  * Your NumberContainers object will be instantiated and called as such:
  * var obj = new NumberContainers()
  * obj.change(index,number)
  * var param_2 = obj.find(number)
  */