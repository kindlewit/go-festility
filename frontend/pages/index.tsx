import Head from "next/head";
import { Fragment } from "react";
import styles from "../styles/Home.module.css";
import SlotCard from "../components/SlotCard";
import { SlotStruct } from "../entities/SlotType";

export default function Home() {
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

  const slotData: SlotStruct = {
    title: "Pierrot the fool",
    id: 123,
    slot_type: "movie",
    schedule_id: "123BAS",
    screen_id: "SERENE",
    duration: 127,
    start_time: new Date(2022, 12, 21, 14, 30, 0, 0).valueOf(),
    // name: "abc",
    movie_id: 2786,
    directors: ["Jean Luc-Godard", "Fellini"],
  };

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
        <SlotCard {...slotData} />
      </div>
    </Fragment>
  );
}
