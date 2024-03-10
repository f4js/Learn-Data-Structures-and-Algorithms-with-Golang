Breadth First Search

### Introduction
1. Used for searching throw the "Graphs".
2. Graphs are created with nodes and edges
3. Nodes are connected together with edges 
4. It works by checking all the nodes that are connected to the main node.
5. It has time notation of O(V+E), V=number of the nodes, E=number of each edge

### Steps to make a Linear Search
1. First create or get a graph
2. Determine what name you are looking at in graph
3. We need a empty searched and queue slice
4. Now we can add the nodes that are connected to us to the queue
5. We then loop throw the queue and get first index of the queue in every loop we must do it
6. We check if this is the name we want or not 
7. If name is not match we will remove the name from the queue and add its connections to queue
8. we do this until the queue is finished, means until we have checked all the nodes


### Pros
1. It's simple
2. Guaranteed to find the shortest path

### Cons
