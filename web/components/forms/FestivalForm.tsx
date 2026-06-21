import React, { Fragment, ReactNode, useActionState } from "react";

import styles from "../../styles/form";
import { generateUniqueFestId } from "../../util/fest";

type FestivalFormProps = {
  children: ReactNode;
  initFormState?: any;
};

export default function FestivalForm({
  children,
  initFormState,
}: FestivalFormProps) {
  const handleCreateFestival = async (prevState: any, formData: any) => {
    const festDict = {
      name: formData.get("name"),
      from_date: new Date(formData.get("from_date")).getTime(),
      to_date: new Date(formData.get("to_date")).getTime(),
      id: generateUniqueFestId(formData.get("name")),
    };

    const res = await fetch("http://localhost:8080/fest", {
      method: "POST",
      body: JSON.stringify(festDict),
    });

    if (res.status === 201) {
      return `Festival ${festDict.name} created`;
    }

    return `Error. Could not create festival ${festDict.name}.`;
  };

  const [formState, formAction, formIsPending] = useActionState(
    handleCreateFestival,
    initFormState ?? null
  );

  return (
    <Fragment>
      <form action={formAction} className={styles.formContainer}>
        {children}
        {formIsPending ? "Submitting..." : formState}
      </form>
    </Fragment>
  );
}
