import java.util.*;

class Solution {
    public List<String> watchedVideosByFriends(
            List<List<String>> watchedVideos,
            int[][] friends,
            int id,
            int level) {

        int n = friends.length;

        // BFS
        Queue<Integer> queue = new LinkedList<>();
        boolean[] visited = new boolean[n];

        queue.offer(id);
        visited[id] = true;

        int currentLevel = 0;

        // Find all friends exactly at the required level
        while (!queue.isEmpty() && currentLevel < level) {
            int size = queue.size();

            for (int i = 0; i < size; i++) {
                int person = queue.poll();

                for (int friend : friends[person]) {
                    if (!visited[friend]) {
                        visited[friend] = true;
                        queue.offer(friend);
                    }
                }
            }

            currentLevel++;
        }

        // Count video frequencies
        Map<String, Integer> frequency = new HashMap<>();

        while (!queue.isEmpty()) {
            int person = queue.poll();

            for (String video : watchedVideos.get(person)) {
                frequency.put(video, frequency.getOrDefault(video, 0) + 1);
            }
        }

        // Put videos into a list
        List<String> result = new ArrayList<>(frequency.keySet());

        // Sort by frequency, then alphabetically
        Collections.sort(result, (a, b) -> {
            if (frequency.get(a) != frequency.get(b)) {
                return frequency.get(a) - frequency.get(b);
            }

            return a.compareTo(b);
        });

        return result;
    }
}