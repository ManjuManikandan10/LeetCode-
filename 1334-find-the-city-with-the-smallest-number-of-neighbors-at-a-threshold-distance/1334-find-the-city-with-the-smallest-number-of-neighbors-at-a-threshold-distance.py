class Solution:
    def findTheCity(self, n: int, edges, distanceThreshold: int) -> int:
        INF = float('inf')

        # Initialize distance matrix
        dist = [[INF] * n for _ in range(n)]

        for i in range(n):
            dist[i][i] = 0

        for u, v, w in edges:
            dist[u][v] = w
            dist[v][u] = w

        # Floyd-Warshall
        for k in range(n):
            for i in range(n):
                for j in range(n):
                    if dist[i][k] + dist[k][j] < dist[i][j]:
                        dist[i][j] = dist[i][k] + dist[k][j]

        answer = -1
        min_reachable = float('inf')

        for i in range(n):
            count = 0

            for j in range(n):
                if i != j and dist[i][j] <= distanceThreshold:
                    count += 1

            # Prefer larger index when counts are equal
            if count <= min_reachable:
                min_reachable = count
                answer = i

        return answer