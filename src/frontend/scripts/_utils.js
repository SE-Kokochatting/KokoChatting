/**
 * description: 工具函数
 * author: Yuming Cui
 * date: 2022-10-11 14:51:36 +0800
 */

import * as child_process from 'child_process'
import * as _readline from 'readline'

/**
 * 执行一段命令行脚本
 * @param {string} command 要执行的 bash 指令
 */
export async function exec(command) {
  return new Promise((resolve, reject) => {
    child_process.exec(command, (err, stdout, stderr) => {
      console.log(stdout)
      if (stderr || err) console.log(err, stderr)
      resolve()
    })
  })
}

/**
 * 打印提示语并从控制台读取一行输入
 * @param {string} hint 展示的提示语
 * @returns 输入的值
 */
export async function readLineSync(hint) {
  const readline = _readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  })
  return new Promise((resolve, reject) => {
    readline.question(hint, (res) => {
      resolve(res)
      readline.close()
    })
  })
}

/**
 * "PostCard" => "post_card"
 * @param {string}} 待转化的字符串
 * @returns 转化后的结果
 */
export const kebabCase = (string) =>
  string
    // 找到小写字母和大写字母的分界处，$1表示匹配到的小写字母，$2表示匹配到的大写字母，在中间插入下划线
    .replace(/([a-z])([A-Z])/g, '$1_$2')
    // 如果存在一到多个空格，将空格替换为下划线
    .replace(/\s+/g, '_')
    .toLowerCase()
