import { FC, ReactElement } from "react";
import dayjs, { Dayjs } from "dayjs";
import utc from "dayjs/plugin/utc";
import timezone from "dayjs/plugin/timezone";

import { DEF_TIMEZONE, SLOT_TIME_FORMAT } from "../constants";
import styles from "../styles/Home.module.css";
import { SlotStruct } from "../entities/SlotType";

dayjs.extend(utc);
dayjs.extend(timezone);

function calculateTop(screenID: number) {
  switch (screenID) {
    case 10: {
      return "30px";
    }
    case 11: {
      return "150px";
    }
    case 12: {
      return "270px";
    }
    case 13: {
      return "390px";
    }
    case 14: {
      return "510px";
    }
    default: {
      return "30px";
    }
  }
}

const SlotCard: FC<SlotStruct> = (slot): ReactElement => {
  let slotStartInIST: Dayjs = dayjs.unix(slot.start_time).tz(DEF_TIMEZONE);
  let slotEndInIST: Dayjs = dayjs
    .unix(slot.start_time + slot.duration * 60)
    .tz(DEF_TIMEZONE);

  let slotPosForCss = {
    x: slotStartInIST.hour() + slotStartInIST.minute() / 60,
    y: calculateTop(slot.screen_id),
  };

  let concatDirectors = slot?.directors?.join(" \u2022 ");

  return (
    <div
      className={styles.slot}
      style={{
        position: "absolute",
        top: slotPosForCss.y,
        left: `calc(${slotPosForCss.x} * var(--time-swimline-width))`,
        width: `calc(${slot.duration / 60} * var(--time-swimline-width))`,
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
        <span className={styles.slot__director}>{slot.screen_name}</span>
        <span className={styles.slot__duration}>{slot.duration}'</span>
      </div>
    </div>
  );
};

export default SlotCard;
