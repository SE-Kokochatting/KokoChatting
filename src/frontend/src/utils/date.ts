
// 2022-11-04T20:04:37+08:00


export function transformTimestamp(timestamp: string) {
  if(timestamp === "" || timestamp === undefined){
    return ""
  }
  let a = new Date(timestamp).getTime();
  const now = new Date()
  const date = new Date(a);
  const Y = date.getFullYear() + '-';
  const M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) ;
  const D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) ;
  const h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) ;
  const m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) ;
  // const s = (date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds()); // 秒
  let dateString = ""
  if(date.getFullYear() === now.getFullYear() && date.getDate() == now.getDate() && date.getMonth() == now.getMonth()){
    dateString += h + ":" + m
  }else{
    dateString += Y + M + "-" + D + " " + h + ":" + m
  }
  // const dateString = Y + M + D + h + m + s;
  // console.log('dateString', dateString); // > dateString 2021-07-06 14:23
  return dateString;
}
