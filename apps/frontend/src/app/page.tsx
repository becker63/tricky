import { config } from '@fortawesome/fontawesome-svg-core'
import '@fortawesome/fontawesome-svg-core/styles.css'
config.autoAddCss = false
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faClone } from '@fortawesome/free-regular-svg-icons'

import css from "styled-jsx/css"
import { Box, Stack, HStack, Flex, styled, Center, StyledComponent } from "styled-system/jsx";
import { Collapsible } from '~/components/ui/collapsible'
import { Button } from '~/components/ui/button'
import { Text } from '~/components/ui/text'
import { Popover, Tooltip } from '@ark-ui/react'
import { faXmark } from '@fortawesome/free-solid-svg-icons'
import { Heading } from '~/components/ui/heading'
import { VT323, TextProps, vt324_base } from './fonts'
import { text } from 'styled-system/recipes'

export const HighLight = styled('div', text, {
  defaultProps: {
    variant: 'heading', style: {
      display: 'inline-block', lineHeight: '0', verticalAlign: 'center', color: "rgb(157,108,132) !important", fontSize: '1.5em',
      ...vt324_base.style
    }
  },

}) as StyledComponent<'div', TextProps>
// TODO: put a blinking arrow in the left hand corner using tooltip label icon telling users to scroll down
export default function Home() {
  return (
    <>
      <VT323 lineHeight={'1.2'} color={'#4a4a4a !important'} letterSpacing={'tighter'} m={'100px'} textAlign={'center'} fontSize={'xx-large'}>
         a corrupted version of De <HighLight>finibus</HighLight> bonorum et malorum, a 1st-century BC text by the Roman statesman and philosopher Cicero, with words altered, added, and removed to make it nonsensical and improper Latin. The first two words themselves are a truncation of dolorem ipsum ("pain itself"). Versions of the Lorem ipsum text have been used in <HighLight>typesetting</HighLight> at leastsince the 1960s, when it was popularized by advertisements for Letraset transfer sheets.[1] Lorem ipsum was introduced to the digital world in the mid-1980s, when Aldus employed it in graphic and <HighLight>word-processing</HighLight> templates for its desktop publishing program PageMaker. Other popular word processors, including Pages and Microsoft Word, have since adopted Lorem ipsum,[2] as have many LaTeX packages,[3][4][5] web content managers such as Joomla! and WordPress, and CSS libraries such as Semantic UI.
      </VT323>
    </>
  );
}
