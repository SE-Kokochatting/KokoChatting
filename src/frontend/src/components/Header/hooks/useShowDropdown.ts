import { useEffect, useState } from 'react'

export const useShowDropDown = () => {
  const [showDropDown, setShowDropDown] = useState(false)

  useEffect(() => {
    if (!showDropDown) return
    const closeDropdown = () => {
      setShowDropDown(false)
    }
    document.addEventListener('click', closeDropdown)
    return () => document.removeEventListener('click', closeDropdown)
  }, [showDropDown])

  return { showDropDown, setShowDropDown }
}
