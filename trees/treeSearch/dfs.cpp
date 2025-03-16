#include <iostream>
#include <vector>
#include <stack>

class Graph
{
private:
    int V;                             // Number of vertices
    std::vector<std::vector<int>> adj; // Adjacency list

public:
    // Constructor
    Graph(int vertices) : V(vertices)
    {
        adj.resize(V);
    }

    // Add an edge to the graph
    void addEdge(int v, int w)
    {
        adj[v].push_back(w); // Add w to v's adjacency list
    }

    // DFS using recursion
    void DFSRecursive(int start)
    {
        // Mark all vertices as not visited
        std::vector<bool> visited(V, false);

        // Call the recursive helper function
        DFSUtil(start, visited);
    }

    // A recursive function used by DFSRecursive
    void DFSUtil(int v, std::vector<bool> &visited)
    {
        // Mark the current node as visited and print it
        visited[v] = true;
        std::cout << v << " ";

        // Recur for all the vertices adjacent to this vertex
        for (int adjacent : adj[v])
        {
            if (!visited[adjacent])
            {
                DFSUtil(adjacent, visited);
            }
        }
    }

    // DFS using stack (iterative approach)
    void DFSIterative(int start)
    {
        // Mark all vertices as not visited
        std::vector<bool> visited(V, false);

        // Create a stack for DFS
        std::stack<int> stack;

        // Push the start vertex
        stack.push(start);

        while (!stack.empty())
        {
            // Pop a vertex from stack
            start = stack.top();
            stack.pop();

            // If not visited, mark it as visited and process it
            if (!visited[start])
            {
                std::cout << start << " ";
                visited[start] = true;
            }

            // Get all adjacent vertices of the popped vertex
            // If an adjacent vertex is not visited, push it to the stack
            for (auto it = adj[start].rbegin(); it != adj[start].rend(); ++it)
            {
                if (!visited[*it])
                {
                    stack.push(*it);
                }
            }
        }
    }
};

int main()
{
    // Create a graph
    Graph g(7); // 7 vertices numbered from 0 to 6

    // Add edges
    g.addEdge(0, 1);
    g.addEdge(0, 2);
    g.addEdge(1, 3);
    g.addEdge(1, 4);
    g.addEdge(2, 5);
    g.addEdge(2, 6);

    std::cout << "DFS starting from vertex 0 (Recursive): ";
    g.DFSRecursive(0);
    std::cout << std::endl;

    std::cout << "DFS starting from vertex 0 (Iterative): ";
    g.DFSIterative(0);
    std::cout << std::endl;

    return 0;
}