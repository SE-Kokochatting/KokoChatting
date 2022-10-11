/**
 * description: tsx、scss模板
 * author: Yuming Cui
 * date: 2022-10-11 14:48:46 +0800
 */

export const tsxTemplate = `
 import "./index.scss";
 // type {COMP_NAME}Props = {
 // };
 function {COMP_NAME}(/* props: {COMP_NAME}Props */) {
   // const {} = props;
   return (
     <div className="{KEBAB_NAME}">
       {COMP_NAME}
     </div>
   );
 }
 export default {COMP_NAME};
 `

export const scssTemplate = `
 .{KEBAB_NAME} {
 }
 `
