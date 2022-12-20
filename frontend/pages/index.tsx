import Head from "next/head";
import { Fragment } from "react";
import styles from "../styles/Home.module.css";
import SlotCard from "../components/SlotCard";
import { SlotStruct } from "../entities/SlotType";
import { GetServerSideProps } from "next";
const { API_ENDPOINT } = process.env;

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

      <h1>Festility</h1>
      <div>
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
    const fid = "Fest2022";

    const res = await fetch(
      `${API_ENDPOINT}/fest/${fid}/schedule?date=2021-12-10`
    );

    if (res.status == 200) {
      const json: [SlotStruct] = await res.json();

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
