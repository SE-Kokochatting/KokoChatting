import React from 'react'
import './index.scss'

interface Props {
  name: string
  style?: React.CSSProperties
  onClick?: () => void
}

function SvgIcon(props: Props) {
  const { name, style, onClick } = props
  const iconName = `#icon-${name}`

  return (
    <svg style={style} onClick={onClick}>
      <use xlinkHref={iconName} />
    </svg>
  )
}

export default SvgIcon
