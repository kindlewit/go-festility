import { GetServerSideProps } from "next";
import { useRouter } from "next/router";

import { FestivalStruct } from "../../entities/FestType";
import { Fragment } from "react";
import Head from "next/head";
import dayjs from "dayjs";

type ServerSideProps = {
  data?: FestivalStruct;
};
type PathProp = {
  fid: string;
};

export default function AdminFestPage(props: ServerSideProps) {
  const festival = props.data;

  if (
    !festival?.id ||
    !festival.name ||
    !(festival.from_date && festival.to_date)
  ) {
    return (
      <Fragment>
        <Head>
          <title>Admin</title>
        </Head>

        <div>
          <h1>Uh oh! There was a problem</h1>
          <p>
            We were unable to fetch the details of this festival. Check the id
            from the URL or try again.
          </p>
        </div>
      </Fragment>
    );
  }

  return (
    <Fragment>
      <Head>
        <title>Admin: {festival.name}</title>
      </Head>

      <div>
        <p>{festival.name}</p>
        <p>{dayjs(festival.from_date).format()}</p>
        <p>{dayjs(festival.to_date).format()}</p>
      </div>
    </Fragment>
  );
}

export const getServerSideProps: GetServerSideProps<ServerSideProps, PathProp> =
  async function (context) {
    const { fid } = context.params!;
    const url = `http://localhost:8080/fest/${fid}`;

    const res = await fetch(url);

    console.log(res.status)

    if (res.status != 200) {
      return {
        notFound: true,
        props: {},
      };
    }

    const json: FestivalStruct = await res.json();

    if (!json || json === null) {
      return {
        notFound: true,
        props: {},
      };
    }

    return {
      props: {
        data: json,
      },
    };
  };
