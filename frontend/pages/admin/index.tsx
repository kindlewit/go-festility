import Head from "next/head";
import { Fragment } from "react";
import FestivalForm from "../../components/forms/FestivalForm";
import dayjs from "dayjs";

import styles from "../../styles/form";

const EXPECTED_FORM_STRING_FORMAT = "YYYY-MM-DD";

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
            <div className="flex flex-row w-full space-between">
              <label htmlFor="name" className="font-semibold text-xl">
                Festival name
              </label>
              <span className="font-light text-red-600"> *</span>
            </div>
            <input
              type="text"
              name="name"
              placeholder="Festival name"
              className={styles.input}
              required={true}
            />

            <div className="flex flex-row w-full space-between">
              <label htmlFor="from_date" className="font-semibold text-xl">
                Start date
              </label>
              <span className="font-light text-red-600"> *</span>
            </div>
            <input
              type="date"
              name="from_date"
              defaultValue={dayjs().format(EXPECTED_FORM_STRING_FORMAT)}
              className={styles.input}
              required={true}
            />

            <div className="flex flex-row w-full space-between">
              <label htmlFor="to_date" className="font-semibold text-xl">
                End date
              </label>
              <span className="font-light text-red-600">*</span>
            </div>
            <input
              type="date"
              name="to_date"
              defaultValue={dayjs().format(EXPECTED_FORM_STRING_FORMAT)}
              className={styles.input}
              required={true}
            />

            <label htmlFor="name" className="font-semibold text-xl">
              Festival website URL
            </label>
            <input
              type="text"
              name="url"
              className={styles.input}
              placeholder="URL for existing festival site"
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
