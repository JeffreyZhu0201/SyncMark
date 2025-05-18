/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-04-17 22:34:17
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-04-17 22:49:29
 * @FilePath: \SyncMark\SyncMark-Extention\entrypoints\content-script\annotation.tsx
 * @Description: File Description Here...
 * 
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved. 
 */
// 简易划词批注内容脚本
function getSelectionText() {
  const sel = window.getSelection();
  return sel && sel.toString();
}

function getPageURL() {
  return window.location.href;
}

function showAnnotationInput(quote: string, range: string) {
  // 简单弹窗，可用更优UI替换
  const text = prompt("添加批注：", "");
  if (text) {
    // 调试输出
    console.log("准备发送批注：", { user: "guest", page_url: getPageURL(), quote, range, text });
    fetch("http://localhost:8080/api/annotation", { // 修正接口路径
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        user: "guest",
        page_url: getPageURL(),
        quote,
        range,
        text,
      }),
    })
    .then(res => res.json())
    .then(data => {
      console.log("批注提交结果：", data);
      alert("批注提交成功！");
    })
    .catch(err => {
      console.error("批注提交失败：", err);
      alert("批注提交失败！");
    });
  }
}

// 高亮选中文本并返回高亮span元素
function highlightSelection() {
  const sel = window.getSelection();
  if (!sel || sel.rangeCount === 0) return null;
  const range = sel.getRangeAt(0);
  if (range.collapsed) return null;

  // 创建高亮span
  const span = document.createElement('span');
  span.style.background = '#ffe066';
  span.style.borderRadius = '2px';
  span.style.padding = '0 2px';
  span.style.cursor = 'pointer';
  span.setAttribute('data-syncmark', 'highlight');

  // 包裹选区
  range.surroundContents(span);
  sel.removeAllRanges();
  return span;
}

// 创建自定义右键菜单
function createContextMenu(x: number, y: number, onClick: () => void) {
  // 移除旧菜单
  const old = document.getElementById('syncmark-context-menu');
  if (old) old.remove();

  const menu = document.createElement('div');
  menu.id = 'syncmark-context-menu';
  menu.style.position = 'fixed';
  menu.style.zIndex = '99999';
  menu.style.top = y + 'px';
  menu.style.left = x + 'px';
  menu.style.background = '#fff';
  menu.style.border = '1px solid #ccc';
  menu.style.borderRadius = '4px';
  menu.style.boxShadow = '0 2px 8px rgba(0,0,0,0.15)';
  menu.style.padding = '6px 16px';
  menu.style.fontSize = '14px';
  menu.style.cursor = 'pointer';
  menu.textContent = '高亮并批注';

  menu.onclick = () => {
    onClick();
    menu.remove();
  };
  document.body.appendChild(menu);

  // 点击其他地方关闭菜单
  setTimeout(() => {
    document.addEventListener('mousedown', function handler(e) {
      if (!menu.contains(e.target as Node)) {
        menu.remove();
        document.removeEventListener('mousedown', handler);
      }
    });
  }, 0);
}

// 右键事件处理
document.addEventListener('contextmenu', (e) => {
  const quote = getSelectionText();
  if (quote && quote.length > 0) {
    e.preventDefault();
    createContextMenu(e.clientX, e.clientY, () => {
      // 高亮
      const span = highlightSelection();
      if (span) {
        // 发送到后端
        showAnnotationInput(quote, ''); // range参数可扩展
      }
    });
  }
});

document.addEventListener("mouseup", () => {
  const quote = getSelectionText();
  if (quote && quote.length > 0) {
    // 可扩展为更优的定位方式
    showAnnotationInput(quote, "");
  }
});

// 展示所有批注（可扩展为更优UI）
fetch("http://localhost:8080/api/annotation?page_url=" + encodeURIComponent(getPageURL()))
  .then(res => res.json())
  .then(list => {
    // 调试输出
    console.log("已加载批注：", list);
    list.forEach((ann: any) => {
      // 可用更优UI替换
      console.log("批注：", ann.quote, ann.text);
    });
  })
  .catch(err => {
    console.error("加载批注失败：", err);
  });
