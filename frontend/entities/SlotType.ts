export type SlotStruct = {
  // name: string;
  id: number;
  // from_date: number;
  // to_date: number;
  slot_type: string;
  schedule_id: string;
  screen_id: string;
  start_time: number;
  duration: number;
  title: string;
  year?: string;
  movie_id: number;
  directors?: string[];
};

export type TempSlotStruct = {
  movie_id: number;
  duration?: number;
  title?: string;
  year?: string;
  screen_id: number;
  slot_type: string;
  start_time: number;
};
