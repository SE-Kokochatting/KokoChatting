import './index.scss'

interface Props {
  name: string
  class?: string
}

function SvgIcon(props: Props) {
  const iconName = `#icon-${props.name}`
  const svgClass = props.class ? 'svg-icon ' + props.class : 'svg-icon'

  return (
    <svg className={svgClass}>
      <use xlinkHref={iconName} />
    </svg>
  )
}

export default SvgIcon
