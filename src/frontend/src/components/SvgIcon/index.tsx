// 使用 emotion.js，在 css 属性中可以支持 '&:hover' 等写法
import { CSSObject } from '@emotion/react'
import './index.scss'

interface Props {
  name: string
  style?: CSSObject
  onClick?: () => void
}

function SvgIcon(props: Props) {
  const { name, style, onClick } = props
  const iconName = `#icon-${name}`

  return (
    <svg css={style} onClick={onClick}>
      <use xlinkHref={iconName} />
    </svg>
  )
}

export default SvgIcon
