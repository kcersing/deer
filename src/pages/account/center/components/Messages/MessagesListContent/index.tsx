import { Modal,Button  } from 'antd';
import dayjs from 'dayjs';
import React, { useState } from 'react';
import useStyles from './index.style';

const MessagesListContent: React.FC<MessagesSend> = ({
  data: { content, fromUserName, receivedAt, type },
}) => {

  function RichText({ content }) {
    return <div dangerouslySetInnerHTML={{ __html: content }} />;
  }
  const [isModalOpen, setIsModalOpen] = useState(false);

  const showModal = () => {
    setIsModalOpen(true);
  };

  const handleOk = () => {
    setIsModalOpen(false);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };
  const { styles } = useStyles();
  return (
    <>

    <div  onClick={showModal}>
      <div className={styles.description}>{ RichText({content})}</div>
      <div className={styles.extra}>
        <a>{fromUserName}</a> 发布于
        <em>{dayjs(receivedAt).format('YYYY-MM-DD HH:mm')}</em>
      </div>
    </div>
      <Modal title="Basic Modal" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
        <div>
          <div className={styles.description}>{ RichText({content})}</div>
          <div className={styles.extra}>
            <a>{fromUserName}</a> 发布于
            <em>{dayjs(receivedAt).format('YYYY-MM-DD HH:mm')}</em>
          </div>
        </div>
      </Modal>
    </>
  );
};
export default MessagesListContent;
