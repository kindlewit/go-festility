import Head from "next/head";
import { Fragment } from "react";


export default function HomePage() {
  return (
    <Fragment>
      <Head>
        <title>Festility</title>
      </Head>

      <div aria-label="container" className="flex w-full h-screen bg-dark-bg-3 text-dark-text-1">
        <div aria-label="hero" className="flex flex-col w-full items-center m-10">
          <h1 className="text-4xl font-bold">Welcome to festility</h1>
        </div>
      </div>
    </Fragment>
  )
}
