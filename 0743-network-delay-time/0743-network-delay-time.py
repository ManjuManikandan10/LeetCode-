import heapq
from collections import defaultdict

class Solution:
    def networkDelayTime(self, times, n, k):

        # Build graph
        graph = defaultdict(list)
        for u, v, w in times:
            graph[u].append((v, w))

        # Min-heap: (time, node)
        heap = [(0, k)]
        dist = {}

        while heap:
            time, node = heapq.heappop(heap)

            if node in dist:
                continue

            dist[node] = time

            for nei, w in graph[node]:
                if nei not in dist:
                    heapq.heappush(heap, (time + w, nei))

        # If not all nodes received signal
        if len(dist) != n:
            return -1

        # Maximum time among all nodes
        return max(dist.values())
