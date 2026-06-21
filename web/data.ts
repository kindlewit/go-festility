import { type } from "os";
import { ScreenStruct } from "./entities/ScreenType";
import type { TempSlotStruct } from "./entities/SlotType";

export const Ciff = {};

export const Screen_List: ScreenStruct[] = [
  {
    cinema_name: "SPI Sathyam Cinemas",
    screen_name: "Sathyam",
    id: 10,
    city: "Chennai",
  },
  {
    cinema_name: "SPI Sathyam Cinemas",
    screen_name: "Santham",
    id: 11,
    city: "Chennai",
  },
  {
    cinema_name: "SPI Sathyam Cinemas",
    screen_name: "Serene",
    id: 12,
    city: "Chennai",
  },
  {
    cinema_name: "SPI Sathyam Cinemas",
    screen_name: "Seasons",
    id: 13,
    city: "Chennai",
  },
  {
    cinema_name: "SPI Sathyam Cinemas",
    screen_name: "6 Degrees",
    id: 14,
    city: "Chennai",
  },
  {
    cinema_name: "Anna Cinemas",
    screen_name: "Anna Cinema",
    id: 20,
    city: "Chennai",
  },
];

export const Ciff_Schedule: TempSlotStruct[] = [
  {
    movie_id: 958496, // imaginary country
    slot_type: "movie",
    screen_id: 11,
    start_time: 1671699600,
  },
  {
    movie_id: 587668, // gravedigger
    screen_id: 11,
    start_time: 1671684300,
    slot_type: "movie",
  },
  {
    movie_id: 793823, // glaszimmer
    screen_id: 11,
    start_time: 1671692400,
    slot_type: "movie",
  },
  {
    movie_id: 974521, // panahi
    slot_type: "movie",
    screen_id: 10,
    start_time: 1671712200,
  },
  {
    movie_id: 836018, // paloma
    slot_type: "movie",
    screen_id: 12,
    start_time: 1671684300,
  },
  {
    movie_id: 824387, // cinema bandi
    slot_type: "movie",
    screen_id: 12,
    start_time: 1671693300,
  },
  {
    movie_id: 828373, // between us,
    slot_type: "movie",
    screen_id: 13,
    start_time: 1671683400,
  },
  {
    movie_id: 788734, // sick of myself
    screen_id: 13,
    start_time: 1671692400,
    slot_type: "movie",
  },
  {
    movie_id: 680071, // vanishing
    screen_id: 13,
    start_time: 1671701400,
    slot_type: "movie",
  },
  {
    movie_id: 486776, // losers adventure
    screen_id: 14,
    start_time: 1671681600,
    slot_type: "movie",
  },
  {
    movie_id: 747311, // sami
    screen_id: 14,
    start_time: 1671691500,
    slot_type: "movie",
  },
];
