import { LikeOutlined, MessageFilled, StarTwoTone } from '@ant-design/icons';
import { useRequest } from '@umijs/max';
import { List, Tag } from 'antd';
import React from 'react';

import useStyles from './index.style';
import {MessagesSend} from "../../service/data";
import {getMessagesSendList} from "../../service/service";
import MessagesListContent from "./MessagesListContent";

const Messages: React.FC = () => {
  const { styles } = useStyles();
  const IconText: React.FC<{
    icon: React.ReactNode;
    text: React.ReactNode;
  }> = ({ icon, text }) => (
    <span>
      {icon} {text}
    </span>
  );

  // 获取tab列表数据
  const { data: listData } = useRequest(() => {
    return getMessagesSendList({
      count: 30,
    });
  });

  return (
    <List<MessagesSend>
      size="large"
      className={styles.articleList}
      rowKey="id"
      itemLayout="vertical"
      dataSource={listData || []}
      style={{
        margin: '0 -24px',
      }}
      renderItem={(item) => (

        <List.Item key={item.id}>
          <List.Item.Meta
            title={
              <a>{item.title }</a>
            }
            description={
              <span>
                <Tag>Ant Design</Tag>
                <Tag>设计语言</Tag>
                <Tag>蚂蚁金服</Tag>
              </span>
            }
          />
          <MessagesListContent data={item} />
        </List.Item>
      )}
    />
  );
};
export default Messages;
