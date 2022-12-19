import { FC, ReactElement } from "react";
import { SlotStruct } from "../entities/SlotType";
import styles from "../styles/Home.module.css";

export function timeToValue(time: Date): number {
  return time.getHours() + time.getMinutes() / 60;
}

const SlotCard: FC<SlotStruct> = (slot): ReactElement => {
  let startTime = new Date(slot.start_time);
  const startTimeAsValue = timeToValue(startTime);
  let directors = slot?.directors?.join(' \u2022 ');
  return (
    <div
      className={styles.slot}
      style={{
        position: "relative",
        top: "-300px",
        left: `calc(${startTimeAsValue} * 124px)`,
        width: `calc(${slot.duration / 60} * 124px)`,
      }}
    >
      <div className={styles.slot_details}>
        <span>{startTime.toDateString()}</span>
        <span>
          <strong>{slot.title}</strong>
        </span>
        <span>{directors}</span>
        <span className={styles.slot__duration}>{slot.duration}'</span>
      </div>
    </div>
  );
};

export default SlotCard;
