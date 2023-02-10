// Edmonds Karp
// O(V * E^2)
#include<bits/stdc++.h>
using namespace std;

const int maxn = 512, maxm = 512;
int n, m; // n is number of vertices and m is number of edges
int s, t; // s is source vertex and t is sink
int par[maxn], par_id[maxn];
bool vis[maxn];
vector < pair < int, int > > adj[maxn];
int capacity[maxm];

void get_input();
int max_flow();
bool bfs(int);
void reset_vis_par();


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
    reset_vis_par();
    int result = 0;
    while(bfs(s)) {
        int now = t, mn = 1e9 + 12;
        while (now != s) {
            mn = min(mn, capacity[par_id[now]]);
            now = par[now];
        }
        now = t;
        while (now != s) {
            capacity[par_id[now]] -= mn;
            capacity[par_id[now]^1] += mn;
            now = par[now];
        }
        result += mn;
        reset_vis_par();
    }
    return result;
}

bool bfs(int root) {
    queue < int > q;
    q.push(s);
    vis[s] = true;
    while (!q.empty()) {
        int now = q.front();
        q.pop();
        for (auto x : adj[now]) {
            if (!vis[x.first] && capacity[x.second] > 0) {
                par[x.first] = now;
                par_id[x.first] = x.second;
                vis[x.first] = true;
                if (x.first == t)
                    return true;
                q.push(x.first);
            }
        }
    }
    return false;
}

void reset_vis_par() {
    par[s] = -1;
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