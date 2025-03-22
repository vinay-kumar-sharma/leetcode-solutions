/**
 * @param {number} n
 * @param {number[][]} edges
 * @return {number}
 */
var countCompleteComponents = function(n, edges) {
    let adj = new Map();
    
    // Initialize adjacency list
    for (let i = 0; i < n; i++) adj.set(i, new Set());
    for (let [u, v] of edges) {
        adj.get(u).add(v);
        adj.get(v).add(u);
    }
    
    let visited = new Set();
    let completeCount = 0;
    
    // Function to perform BFS and check completeness
    function bfs(node) {
        let queue = [node];
        let component = new Set([node]);
        visited.add(node);
        
        while (queue.length > 0) {
            let curr = queue.shift();
            for (let neighbor of adj.get(curr)) {
                if (!visited.has(neighbor)) {
                    visited.add(neighbor);
                    queue.push(neighbor);
                    component.add(neighbor);
                }
            }
        }
        
        // Check if this component is complete
        let size = component.size;
        let edgeCount = 0;
        for (let node of component) {
            edgeCount += adj.get(node).size;
        }
        edgeCount /= 2;  // Each edge is counted twice
        
        return edgeCount === (size * (size - 1)) / 2; // Check completeness
    }
    
    // Traverse all nodes
    for (let i = 0; i < n; i++) {
        if (!visited.has(i)) {
            if (bfs(i)) completeCount++;
        }
    }
    
    return completeCount;
};
