# 数据结构

## 数据结构演进

数据结构的演进体现了从简单到复杂、从线性到网状的思维转变:

- **链表 → 二叉树**: 从线性思维到递归思维，每个节点从最多 1 个子节点扩展到 2 个
- **二叉树 → 树**: 从递归到分治，节点可以有任意多个子节点
- **树 → 图**: 从层次结构到网状结构，打破了"无环"的限制，节点间可以有任意连接

$$
\begin{array}{c c c c c c c}
  \textbf{链表}
  & \xtofrom[\text{每个节点最多1个子节点}]{\text{线性}\to\text{递归思维}}
  & \textbf{二叉树}
  & \xtofrom[\text{每个节点最多2个子节点}]{\text{分治思维}}
  & \textbf{树}
  & \xtofrom[\text{无环且连通}]{\text{网状思维}}
  & \textbf{图} \\
  \begin{matrix} \scriptsize\text{一对一} \\ \scriptsize\text{线性结构} \end{matrix} &&
  \begin{matrix} \scriptsize\text{一对二} \\ \scriptsize\text{层次结构} \end{matrix} &&
  \begin{matrix} \scriptsize\text{一对多} \\ \scriptsize\text{层次结构} \end{matrix} &&
  \begin{matrix} \scriptsize\text{多对多} \\ \scriptsize\text{网状结构} \end{matrix} &&
\end{array}
$$
