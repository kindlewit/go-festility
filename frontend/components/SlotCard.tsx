import { FC, ReactElement } from "react";
import dayjs, { Dayjs } from "dayjs";
import utc from "dayjs/plugin/utc";
import timezone from "dayjs/plugin/timezone";

import { DEF_TIMEZONE, SLOT_TIME_FORMAT } from "../constants";
import styles from "../styles/Home.module.css";
import { SlotStruct } from "../entities/SlotType";

dayjs.extend(utc);
dayjs.extend(timezone);

export function timeToValue(time: Dayjs): number {
  return time.hour() + time.minute() / 60;
}

const SlotCard: FC<SlotStruct> = (slot): ReactElement => {
  let slotStartInIST: Dayjs = dayjs.unix(slot.start_time).tz(DEF_TIMEZONE);
  let slotEndInIST: Dayjs = dayjs
    .unix(slot.start_time + slot.duration * 60)
    .tz(DEF_TIMEZONE);

  let slotPosForCss = {
    x: slotStartInIST.hour() + slotStartInIST.minute() / 60,
  };

  let concatDirectors = slot?.directors?.join(" \u2022 ");

  return (
    <div
      className={styles.slot}
      style={{
        position: "relative",
        top: "-500px",
        left: `calc(${slotPosForCss.x} * 124px)`,
        width: `calc(${slot.duration / 60} * 124px)`,
      }}
    >
      <div className={styles.slot_details}>
        <span>
          {slotStartInIST.format(SLOT_TIME_FORMAT)} –{" "}
          {slotEndInIST.format(SLOT_TIME_FORMAT)}
        </span>
        <span>
          <strong>{slot.title}</strong>
        </span>
        <span className={styles.slot__director}>{concatDirectors}</span>
        <span className={styles.slot__duration}>{slot.duration}'</span>
      </div>
    </div>
  );
};

export default SlotCard;
