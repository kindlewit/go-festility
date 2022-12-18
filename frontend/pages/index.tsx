import Head from "next/head";
import Image from "next/image";
import { Fragment } from "react";
import styles from "../styles/Home.module.css";

export default function Home() {
  return (
    <Fragment>
      <Head>
        <title>Festility</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <div>
        <div style={{ display: "flex", flexDirection: "column" }}>
          <div style={{ display: "flex", flexDirection: "row" }}>
            <div className={styles.flex_item}>1</div>
            <div className={styles.flex_item}>2</div>
            <div className={styles.flex_item}>3</div>
            <div className={styles.flex_item}>4</div>
            <div className={styles.flex_item}>5</div>
            <div className={styles.flex_item}>6</div>
            <div className={styles.flex_item}>7</div>
            <div className={styles.flex_item}>8</div>
            <div className={styles.flex_item}>9</div>
            <div className={styles.flex_item}>10</div>
            <div className={styles.flex_item}>11</div>
            <div className={styles.flex_item}>12</div>
            <div className={styles.flex_item}>13</div>
            <div className={styles.flex_item}>14</div>
            <div className={styles.flex_item}>15</div>
            <div className={styles.flex_item}>16</div>
            <div className={styles.flex_item}>17</div>
            <div className={styles.flex_item}>18</div>
            <div className={styles.flex_item}>19</div>
            <div className={styles.flex_item}>20</div>
            <div className={styles.flex_item}>21</div>
            <div className={styles.flex_item}>22</div>
            <div className={styles.flex_item}>23</div>
            <div className={styles.flex_item}>24</div>
          </div>
          <div style={{ display: "flex", flexDirection: "row" }}>
            <div className={styles.flex_item}>1</div>
            <div className={styles.flex_item}>2</div>
            <div className={styles.flex_item}>3</div>
            <div className={styles.flex_item}>4</div>
            <div className={styles.flex_item}>5</div>
            <div className={styles.flex_item}>6</div>
            <div className={styles.flex_item}>7</div>
            <div className={styles.flex_item}>8</div>
            <div className={styles.flex_item}>9</div>
            <div className={styles.flex_item}>10</div>
            <div className={styles.flex_item}>11</div>
            <div className={styles.flex_item}>12</div>
            <div className={styles.flex_item}>13</div>
            <div className={styles.flex_item}>14</div>
            <div className={styles.flex_item}>15</div>
            <div className={styles.flex_item}>16</div>
            <div className={styles.flex_item}>17</div>
            <div className={styles.flex_item}>18</div>
            <div className={styles.flex_item}>19</div>
            <div className={styles.flex_item}>20</div>
            <div className={styles.flex_item}>21</div>
            <div className={styles.flex_item}>22</div>
            <div className={styles.flex_item}>23</div>
            <div className={styles.flex_item}>24</div>
          </div>
        </div>
      </div>
    </Fragment>
  );
}
