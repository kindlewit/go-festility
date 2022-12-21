// Temp fetch schedule data
import type { NextApiRequest, NextApiResponse } from "next";

import type { TempSlotStruct } from "../../../entities/SlotType";
import { Ciff_Schedule } from "../../../data";

const { MOVIE_API_KEY } = process.env;

async function fetchMovieData(slot: TempSlotStruct) {
  const url = `https://api.themoviedb.org/3/movie/${slot.movie_id}?api_key=${MOVIE_API_KEY}`;
  const res = await fetch(url);
  const json = await res.json();

  return {
    ...slot,
    title: json.title,
    duration: json.runtime,
    year: json.release_date.match(/\d{4}/ig)[0],
  }
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<TempSlotStruct[]>
) {
  let data = await Promise.all(Ciff_Schedule.map(fetchMovieData));
  res.status(200).json(data);
}
