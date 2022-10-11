import React from 'react'
import './index.scss'

interface Props {
  name: string
  style?: React.CSSProperties
}

function SvgIcon(props: Props) {
  const { name, style } = props
  const iconName = `#icon-${name}`

  return (
    <svg style={style}>
      <use xlinkHref={iconName} />
    </svg>
  )
}

export default SvgIcon
