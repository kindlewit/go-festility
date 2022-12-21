import Head from "next/head";
import { Fragment } from "react";
import { GetServerSideProps } from "next";

import SlotCard from "../components/SlotCard";
import { SlotStruct } from "../entities/SlotType";
import styles from "../styles/Home.module.css";

// const { API_ENDPOINT } = process.env;

type ServerSideProps = {
  data?: SlotStruct[];
};

type PathProp = {
  fid: string;
};

export default function Home(props: { data: SlotStruct[] }) {
  const hoursDisplay = Array.from({ length: 24 }).map(
    (v: unknown, i: number) => {
      // if (i < 9) {
      //   // Temporarily no slots before 9 am
      //   return (
      //     <div key={i} className={styles.flex_item}>
      //       {i + ":00"}
      //     </div>
      //   );
      // }
      return (
        <div key={i} className={styles.flex_item}>
          {i + ":00"}
        </div>
      );
    }
  );

  let slotData = props.data.map((d) => (
    <section>
      <SlotCard {...d} />
    </section>
  ));

  return (
    <Fragment>
      <Head>
        <title>Festility</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <div style={{ position: "relative" }}>
        <div style={{ display: "flex", flexDirection: "column" }}>
          <div style={{ display: "flex", flexDirection: "row" }}>
            {hoursDisplay}
          </div>
        </div>
        {slotData}
      </div>
    </Fragment>
  );
}

export const getServerSideProps: GetServerSideProps<ServerSideProps, PathProp> =
  async function (context) {
    // const fid = "Fest2022";
    // const url = `${API_ENDPOINT}/fest/${fid}/schedule?date=2021-12-10`;
    const url = "http://localhost:3000/api/fest/1";

    const res = await fetch(url);

    if (res.status == 200) {
      const json: SlotStruct[] = await res.json();

      if (json != null) {
        return {
          props: {
            data: json,
          },
        };
      }
    }
    return {
      notFound: true,
      props: {},
    };
  };
