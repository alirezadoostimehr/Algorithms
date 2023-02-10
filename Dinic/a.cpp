// Dinic
// O(V^3)
#include <bits/stdc++.h>
using namespace std;

const int maxn = 512, maxm = 512;
int n, m; // n is number of vertices and m is number of edges
int s, t; // s is source vertex and t is sink
int level[maxn], ind[maxn];
bool vis[maxn];
vector < pair < int, int > > adj[maxn];
int capacity[maxm];

void get_input();
int max_flow();
bool bfs(int);
void reset_vis();
void reset_level();
void reset_ind();
int dfs(int, int);


int main() {
    get_input();
    cout << "max possible flow: " << max_flow();
    return 0;
}

void get_input() {
    cin >> n >> m;
    cin >> s >> t;
    s--; t--;
    int id = 0;
    for (int i = 0; i < m; i++) {
        int u, v, c;
        cin >> u >> v >> c;
        u--; v--;
        adj[u].push_back(make_pair(v, id));
        adj[v].push_back(make_pair(u, id^1));
        capacity[id] = c;
        id += 2;
    }
    return;
}

int max_flow() {
    reset_level();
    int result = 0;
    while(bfs(s)) {
        reset_ind();
        result += dfs(s, INT_MAX);
    }
    return result;
}

int dfs(int root, int input_flow) {
    if (root == t)
        return input_flow;
    int result = 0;
    for (; ind[root] < adj[root].size(); ind[root]++) {
        auto x = adj[root][ind[root]];
        if ((capacity[x.second] > 0) && (level[x.first] == (level[root] + 1))) {
            int tmp = dfs(x.first, min(input_flow, capacity[x.second]));
            input_flow -= tmp;
            capacity[x.second] -= tmp;
            capacity[x.second^1] += tmp;
            result += tmp;
        }
    }
    return result;
}

void reset_ind() {
    for (int i = 0; i < maxn; i++)
        ind[i] = 0;
    return;
}

bool bfs(int root) {
    queue < int > q;
    q.push(s);
    for (int i = 0; i < n; i++)
        level[i] = -1;
    level[s] = 0;
    while (!q.empty()) {
        int now = q.front();
        q.pop();
        for (auto x : adj[now]) {
            if (level[x.first] < 0 && capacity[x.second] > 0) {
                level[x.first] = level[now] + 1;
                q.push(x.first);
            }
        }
    }
    return (level[t] != -1);
}

void reset_level() {
    for (int i = 0; i < n; i++)
        level[i] = -1;
    return;
}

void reset_vis() {
    for (int i = 0; i < n; i++)
        vis[i] = false;
    return;
}

/*
# Sample Input
6 10
1 6
1 2 16
1 3 13
2 3 10
2 4 12
3 2 4
3 5 14
4 3 9
4 6 20
5 4 7
5 6 4

# Sample Output
23
*/
