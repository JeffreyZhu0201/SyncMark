/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-16 13:40:35
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-04-17 22:52:45
 * @FilePath: \SyncMark\SyncMark-Extention\wxt.config.ts
 * @Description: File Description Here...
 * 
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved. 
 */
import { defineConfig } from 'wxt';

// See https://wxt.dev/api/config.html
export default defineConfig({
  extensionApi: 'chrome',
  modules: ['@wxt-dev/module-react']
});
