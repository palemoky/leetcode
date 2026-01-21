# 排序

排序算法是算法与数据结构领域的基础内容，广泛应用于数据处理、查找、去重等场景。掌握常见排序算法及其原理、复杂度和适用场景，是算法学习和面试的必备技能。

排序的方法有 **插入**、**交换**、**选择**、**合并** 等。

## 十大常用基础排序算法

<table>
  <thead>
    <tr>
      <th>名称</th>
      <th>数据对象</th>
      <th>稳定性</th>
      <th>比较类</th>
      <th>时间复杂度（平均/最坏）</th>
      <th>空间复杂度</th>
      <th>原理</th>
      <th>描述</th>
      <th>适用场景</th>
      <th>动画</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><strong>冒泡排序</strong><br />（bubble sort）</td>
      <td>数组</td>
      <td>✓</td>
      <td>✓</td>
      <td markdown="1">$O(n^2)$</td>
      <td markdown="1">$O(1)$</td>
      <td>每轮将相邻元素两两比较，大的往后交换，重复 n 轮</td>
      <td>(无序区, 有序区)。<br />从无序区通过交换找出最大元素放到有序区前端。</td>
      <td>数据量小、对稳定性有要求</td>
      <td><a href="https://www.bilibili.com/video/BV181421876R">冒泡排序</a></td>
    </tr>
    <tr>
      <td rowspan="2" style="vertical-align: middle;"><strong>选择排序</strong><br />（selection sort）</td>
      <td>数组</td>
      <td>×</td>
      <td rowspan="2" style="vertical-align: middle;">✓</td>
      <td rowspan="2" style="vertical-align: middle;" markdown="1">$O(n^2)$</td>
      <td rowspan="2" style="vertical-align: middle;" markdown="1">$O(1)$</td>
      <td rowspan="2" style="vertical-align: middle;">每轮选择剩余元素中的最小值，放到前面</td>
      <td rowspan="2" style="vertical-align: middle;">(有序区, 无序区)。<br />在无序区里找一个最小的元素跟在有序区的后面。对数组：比较得多，换得少。</td>
      <td rowspan="2" style="vertical-align: middle;">数据量小</td>
      <td rowspan="2" style="vertical-align: middle;"><a href="https://www.bilibili.com/video/BV1kjsuenE8v">选择排序</a></td>
    </tr>
    <tr>
      <td>链表</td>
      <td>✓</td>
    </tr>
    <tr>
      <td><strong>插入排序</strong><br />（insertion sort）</td>
      <td>数组、链表</td>
      <td>✓</td>
      <td>✓</td>
      <td markdown="1">$O(n^2)$</td>
      <td markdown="1">$O(1)$</td>
      <td>每次将一个元素插入到已排序部分的合适位置</td>
      <td>(有序区, 无序区)。<br />把无序区的第一个元素插入到有序区的合适位置。对数组：比较得少，换得多。</td>
      <td>数据量小、部分有序</td>
      <td></td>
    </tr>
    <tr>
      <td><strong>堆排序</strong><br />（heap sort）</td>
      <td>数组</td>
      <td>×</td>
      <td>✓</td>
      <td markdown="1">$O(n \log n)$</td>
      <td markdown="1">$O(1)$</td>
      <td>构建最大/最小堆，依次取出堆顶元素</td>
      <td>(最大堆, 有序区)。<br />从堆顶把根卸出来放在有序区之前，再恢复堆。</td>
      <td>原地排序</td>
      <td><a href="https://www.bilibili.com/video/BV1HYtseiEQ8">堆排序</a></td>
    </tr>
    <tr>
      <td rowspan="2" style="vertical-align: middle;"><strong>归并排序</strong><br />（merge sort）</td>
      <td>数组</td>
      <td rowspan="2" style="vertical-align: middle;">✓</td>
      <td rowspan="2" style="vertical-align: middle;">✓</td>
      <td markdown="1">$O(n \log n)$</td>
      <td markdown="1">$O(n) + O(\log n)$</td>
      <td rowspan="2" style="vertical-align: middle;">递归分组，合并有序子数组</td>
      <td rowspan="2" style="vertical-align: middle;">把数据分为两段，从两段中逐个选最小的元素移入新数据段的末尾。可从上到下或从下到上进行。</td>
      <td rowspan="2" style="vertical-align: middle;">大数据、链表排序、稳定性要求高</td>
      <td rowspan="2" style="vertical-align: middle;"><a href="https://www.bilibili.com/video/BV1em1oYTEFf">归并排序</a></td>
    </tr>
    <tr>
      <td>链表</td>
      <td markdown="1">$O(n \log n)$</td>
      <td markdown="1">$O(1)$</td>
    </tr>
    <tr>
      <td rowspan="2" style="vertical-align: middle;"><strong>快速排序</strong><br />（quick sort）</td>
      <td>数组</td>
      <td>×</td>
      <td rowspan="2" style="vertical-align: middle;">✓</td>
      <td rowspan="2" style="vertical-align: middle;" markdown="1">$O(n \log n) / O(n^2)$</td>
      <td rowspan="2" style="vertical-align: middle;" markdown="1">$O(\log n)$</td>
      <td rowspan="2" style="vertical-align: middle;">选定基准，分区递归排序左右两部分</td>
      <td rowspan="2" style="vertical-align: middle;">(小数, 基准元素, 大数)。<br />在区间中随机挑选一个元素作基准，将小于基准的元素放在基准之前，大于基准的元素放在基准之后，再分别对小数区与大数区进行排序。</td>
      <td rowspan="2" style="vertical-align: middle;">通用、高效排序</td>
      <td rowspan="2" style="vertical-align: middle;"><a href="https://www.bilibili.com/video/BV1y4421Z7hK">快速排序</a></td>
    </tr>
    <tr>
      <td>链表</td>
      <td>✓</td>
    </tr>
    <tr>
      <td><strong>希尔排序</strong><br />（shell sort）</td>
      <td>数组</td>
      <td>×</td>
      <td>✓</td>
      <td markdown="1">$O(n \log^2 n) / O(n^2)$</td>
      <td markdown="1">$O(1)$</td>
      <td></td>
      <td>每一轮按照事先决定的间隔进行插入排序，间隔会依次缩小，最后一次一定要是 1。</td>
      <td></td>
      <td></td>
    </tr>
    <tr>
      <td><strong>计数排序</strong><br />（counting sort）</td>
      <td rowspan="3" style="vertical-align: middle;">数组、链表</td>
      <td rowspan="3" style="vertical-align: middle;">✓</td>
      <td rowspan="3" style="vertical-align: middle;">×</td>
      <td markdown="1">$O(n + m)$</td>
      <td markdown="1">$O(n + m)$</td>
      <td rowspan="3" style="vertical-align: middle;">利用元素值域特性进行分组计数或分桶</td>
      <td>统计小于等于该元素的值的元素的个数 i，于是该元素就放在目标数组的索引 i 位 (i≥0)。</td>
      <td rowspan="3" style="vertical-align: middle;">数据范围有限、整数排序</td>
      <td></td>
    </tr>
    <tr>
      <td><strong>桶排序</strong><br />（bucket sort）</td>
      <td markdown="1">$O(n)$ / $O(n^2)$</td>
      <td markdown="1">$O(m)$</td>
      <td>将值为 i 的元素放入 i 号桶，最后依次把桶里的元素倒出来。</td>
      <td></td>
    </tr>
    <tr>
      <td><strong>基数排序</strong><br />（radix sort）</td>
      <td markdown="1">$O(k \times n) / O(n^2)$</td>
      <td markdown="1">$O(n)$</td>
      <td>一种多关键字的排序算法，可用桶排序实现。</td>
      <td><a href="https://www.bilibili.com/video/BV1KrzrYeEDw">基数排序</a></td>
    </tr>
  </tbody>
</table>

!!! Info "表格说明"

    - **n**：数据规模
    - **m**：数据的最大值减最小值
    - **k**：数值中的"数位"个数
    - 所有排序均按从小到大排列

!!! Note "重要说明"

    - **比较类 vs 非比较类**：计数排序、桶排序、基数排序均为非比较类排序。现代编程语言的内置排序（如 C++、Java、Python）都是**比较类排序**（一般 $O(n \log n)$），因为要能支持**通用对象排序**。
    - **必须掌握**：排序算法是算法基础，建议至少熟练掌握冒泡、插入、选择、快排、归并五种实现。
    - **选择依据**：选择合适的排序算法需结合数据规模、稳定性需求和空间限制。

<div align="center">
  <table border="0" cellpadding="0" cellspacing="0">
    <tr>
      <td align="center" style="vertical-align: bottom;">
        <img src="quick_sorting.webp" alt="Quick Sorting" /><br />
        <sub class="img-caption">快速排序示意图</sub>
      </td>
      <td align="center" style="vertical-align: bottom;">
        <img src="merge_sorting.webp" alt="Merge Sorting" /><br />
        <sub class="img-caption">归并排序示意图</sub>
      </td>
    </tr>
  </table>
</div>

!!! Tip "快速记忆口诀"

    - 冒泡两两换，选择找最小，
    - 插入往前挪，希尔改插入。
    - 归并分两半，快排选基准，
    - 堆排建大堆，计数靠统计。
    - 桶排序分区，基数按位排。

### 复杂度

在计算机科学中，我们通常用大 $O$ 来描述某个特定算法时间与空间随着数据规模增加而变化的趋势。

<figure>
    <img src="Big_O.webp" alt="Big-O Complexity Chart" width="50%" />
</figure>

### 稳定性与不稳定性

我们以纸牌排序为例，当纸牌用稳定排序按点值排序的时候，两个 5 之间必定保持它们最初的次序。在用不稳定排序来排序的时候，两个 5 可能被按相反次序来排序。

<div align="center">
  <table border="0" cellpadding="0" cellspacing="0">
    <tr>
      <td align="center" style="vertical-align: bottom;">
        <img src="sorting_stable.webp" alt="稳定排序" width="70%" /><br />
      </td>
      <td align="center" style="vertical-align: bottom;">
        <img src="sorting_unstable.webp" alt="不稳定排序" width="70%" /><br />
      </td>
    </tr>
  </table>
</div>

## 工程常用算法

### Tim 排序（归并+插入）

Timsort 是一中混合（归并+插入）稳定的排序算法。具有 _O_(_n_ log _n_) 的平均和最坏时间复杂度，最优可达 _O_(n)，空间复杂度为 _O_(n)。该算法是目前已知最快的排序算法，在 Python、Swift、Rust 等语言的内置排序功能中被用作默认算法。

### 内省排序（Introsort）（快排+堆排）

内省排序首先从快速排序开始，当递归深度超过一定深度（深度为排序元素数量的对数值）后转为堆排序。采用这个方法，内省排序既能在常规数据集上实现快速排序的高性能，又能在最坏情况下仍保持 _O_(_n_ log _n_) 的时间复杂度。由于这两种算法都属于比较排序算法，所以内省排序也是一个比较排序算法。

### 不同语言内置排序算法对比

| **语言 / 库**                          | **算法实现**                                  | **稳定性** | **特点**                                           |
| -------------------------------------- | --------------------------------------------- | ---------- | -------------------------------------------------- |
| **C (qsort)**                          | 快速排序为主（不同实现可能混合插入排序）      | ×          | 简单高效，但可能退化到 $O(n²)$                     |
| **C++ (std::sort)**                    | **Introsort**（快速排序 + 堆排序 + 插入排序） | ×          | 平均 $O(n log n)$，最坏 $O(n log n)$，避免快排退化 |
| **C++ (std::stable_sort)**             | 归并排序（常带优化）                          | ✓          | 保证稳定性，但需要额外内存                         |
| **Java (Arrays.sort, 基本类型)**       | **Dual-Pivot QuickSort**（双轴快排）          | ×          | 比普通快排常数更小，性能优越                       |
| **Java (Arrays.sort, 对象类型)**       | **Timsort**（归并 + 插入）                    | ✓          | 对部分有序数据非常快，最坏 $O(n log n)$            |
| **Python (list.sort / sorted)**        | **Timsort**                                   | ✓          | 专门为现实数据优化，利用已有有序片段               |
| **JavaScript (V8 引擎)**               | 小数组：插入排序；大数组：快排 + 混合         | ×          | 对象数组时可能使用归并变体                         |
| **JavaScript (SpiderMonkey, Firefox)** | **Timsort / 归并变体**                        | ✓          | 性能和 Python 类似                                 |
| **Go (sort.Sort)**                     | **Introsort**（快排 + 堆排 + 插入）           | ×          | 类似 C++ std::sort                                 |
| **Rust (sort)**                        | **Introsort**（快排 + 堆排 + 插入）           | ×          | 与 C++ 类似                                        |
| **Rust (sort_stable)**                 | **归并排序**                                  | ✓          | 保证稳定性                                         |

- 几乎所有标准库排序都是 **混合算法**，避免单一算法的缺陷，保证最坏复杂度不超过 $O(n log n)$
- 数值数组：大多用快速排序或 Introsort（追求性能）。
- 对象数组（需要稳定性）：多数用 Timsort / 归并排序。
- 小规模数据：经常用 插入排序 优化。
