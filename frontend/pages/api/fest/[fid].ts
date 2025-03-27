// Temp fetch schedule data
import type { NextApiRequest, NextApiResponse } from "next";

import type { TempSlotStruct } from "../../../entities/SlotType";
import { Ciff_Schedule, Screen_List } from "../../../data";

const { MOVIE_API_KEY } = process.env;

async function fetchMovieData(slot: TempSlotStruct) {
  if (
    slot.slot_type == "movie" &&
    (slot.title == undefined || slot.title != null || slot.title != "")
  ) {
    const url = `https://api.themoviedb.org/3/movie/${slot.movie_id}?api_key=${MOVIE_API_KEY}`;
    const res = await fetch(url);
    const json = await res.json();

    return {
      ...slot,
      title: json.title,
      duration: json.runtime,
      year: json.release_date?.match(/\d{4}/gi) && json.release_date?.match(/\d{4}/gi)[0],
    };
  }
  return slot;
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<TempSlotStruct[]>
) {
  let data = await Promise.all(Ciff_Schedule.map(fetchMovieData));

  // Bind screen details
  let boundData = data.map((d) => ({
    ...d,
    ...Screen_List.find((s) => s.id == d.screen_id),
  }));
  // TODO: Group by screen id
  res.status(200).json(boundData);
}
