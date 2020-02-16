import * as React from 'react'
import Head from 'next/head'
import App from 'next/app'
import { withApollo } from '../lib/withApollo';


class MyApp extends App {
  // Only uncomment this method if you have blocking data requirements for
  // every single page in your application. This disables the ability to
  // perform automatic static optimization, causing every page in your app to
  // be server-side rendered.
  //
  // static async getInitialProps(appContext) {
  //   // calls page's `getInitialProps` and fills `appProps.pageProps`
  //   const appProps = await App.getInitialProps(appContext);
  //
  //   return { ...appProps }
  // }

  render() {
    const { Component,  pageProps } = this.props
    return (
      <>
        <Head>
          <title></title>
        </Head>
        <div>
          <Component {...pageProps}/>
        </div>
      </>
    )
  }
}

export default withApollo(MyApp);