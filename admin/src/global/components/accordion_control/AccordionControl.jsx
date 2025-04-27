import { Accordion, Center } from "@mantine/core";
import styles from "./AccordionControl.module.css";
import CustomActionButton from "./CustomActionIcon";


function AccordionControl({ icon, onClick, children, ...props }) {
  return (
    <Center>
      <Accordion.Control {...props} className={styles.accordion_control}>
        <div className={styles.accordion_title_content}>
          {children}

          {icon && onClick && (
            <CustomActionButton
              icon={icon}
              onClick={onClick}
              data-accordion={children}
            />
          )}
        </div>
      </Accordion.Control>
    </Center>
  );
}


export default AccordionControl;