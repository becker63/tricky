import { DotGothic16, VT323 as VT323google } from "next/font/google"
import LocalFont from "next/font/local"
import { ReactElement, ElementType } from "react"
import { styled } from 'styled-system/jsx'
import { type TextVariantProps, text } from 'styled-system/recipes'
import type { ComponentProps, StyledComponent } from 'styled-system/types'

export type TextProps = TextVariantProps & { as?: ElementType }


// Title font
export const dotGothic_base = DotGothic16({
  weight: "400",
  subsets: ['latin'],
})

// accent font
// https://www.dafont.com/silkscreen.font
export const silkscreen_base = LocalFont({
  src: [
    {
      path: "../../public/fonts/slkscre.ttf",
      weight: '400'
    }
  ]
})

// text font
export const vt324_base = VT323google({
  weight: "400",
  subsets: ['latin'],
})



// according to https://github.com/DavidBuchanan314/bitmap-font-css these should help with our font rendering 
// (I honestly cant see a difference but whatever)
const bitmapHelpers = {
  fontSmooth: 'never',
      textSizeAdjust: 'none',
      WebkitFontSmoothing: 'none',
      MozOsxFontSmoothing: 'grayscale',
      filter: 'contrast(100.00001%)',
      color: 'inherit',
}

export const Silkscreen = styled('h2', text, {
  defaultProps: {
    variant: 'heading', style: {
      ...bitmapHelpers,
      ...silkscreen_base.style
    }
  },

}) as StyledComponent<'h2', TextProps>

export const DotGothic = styled('h1', text, {
  defaultProps: {
    variant: 'heading', style: {
      ...bitmapHelpers,
      ...dotGothic_base.style
    }
  },

}) as StyledComponent<'h1', TextProps>

export const VT323 = styled('div', text, {
  defaultProps: {
    variant: 'heading', style: {
      ...bitmapHelpers,
      ...vt324_base.style
    }
  },

}) as StyledComponent<'div', TextProps>

type KernOptions = { right?: string, left?: string }
interface KernSettingsType {
  [key: string]: KernOptions
}

// TODO: Better type for this, less undefined value is dependent on letter
const KernSettings: KernSettingsType = {
  
}

export const DotGothicKerned = ({ children }: { children: string }): ReactElement => {
  return (
    <div>
      {children.split("").map((char) => {
        
        const isInKern = Object.keys(KernSettings).find((k) => k == char) != undefined

        if (isInKern){
          const foundKern = KernSettings[char]
          console.log(char, foundKern)
          return (
            // BUG for some reason margin left as a jsx prop does not work, so we use style here
            <styled.span
              marginRight={foundKern.right ? foundKern.right : ""}
              style={{ marginLeft: foundKern.left ? foundKern.left : "" }}
              className={dotGothic_base.className}
            >
              {char}
            </styled.span>
          )
        }
        return <span className={dotGothic_base.className}>{char}</span>
      })}
    </div>
  )
}