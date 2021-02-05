import React from 'react'
import './App.css'
import InlineSVG from 'svg-inline-react'

import perspective from './images/senserver_perspective.svg'

import { css } from 'pretty-lights'

const headerImage = css`
  height: 300px;
`

function App() {
  return (
    <div className='App'>
      <div className={headerImage}>
        <InlineSVG raw={true} src={perspective} />
      </div>
    </div>
  )
}

export default App
