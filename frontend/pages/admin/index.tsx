import Head from "next/head";
import { Fragment } from "react";
import FestivalForm from "../../components/forms/FestivalForm";

import styles from "../../styles/form";

export default function AdminIndex() {
  return (
    <Fragment>
      <Head>
        <title>Festility admin</title>
      </Head>

      <div
        aria-label="container"
        className="flex flex-col w-full h-screen bg-light-bg-1"
      >
        <h2 className="text-3xl font-bold text-light-text-1">
          Create a festival
        </h2>

        <div>
          <FestivalForm>
            <label htmlFor="name" className="font-semibold text-xl">
              Festival name
            </label>
            <input
              type="text"
              name="name"
              placeholder="Festival name"
              className={styles.input}
            />
            <label htmlFor="name" className="font-semibold text-xl">
              Start date
            </label>
            <input
              type="date"
              name="from_date"
              defaultValue={new Date().getTime()}
              className={styles.input}
            />
            <label htmlFor="name" className="font-semibold text-xl">
              End date
            </label>
            <input
              type="date"
              name="to_date"
              defaultValue={new Date().getTime()}
              className={styles.input}
            />
            <button
              type="submit"
              className="p-3 my-3 border-none rounded-lg bg-light-accent-1 text-white"
            >
              CREATE
            </button>
          </FestivalForm>
        </div>
      </div>
    </Fragment>
  );
}
