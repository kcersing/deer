import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm,
  ProForm,
  ProFormSelect,
  ProFormText,
  ProFormDateRangePicker,
  ProFormTextArea,
  ProFormUploadButton,ProFormDigit,
  ProFormSwitch,
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message } from 'antd';
import React, { FC,useState } from 'react';

import { Messages } from  "./../service/data";
import {getMessagesTypes,SendMessages} from "./../service/service";
interface CreateFormProps {
  reload?: ActionType['reload'];
}
import WangEditor from '@/pages/components/wangeditor'


const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();


  const [type, setType] = useState("card");


  const { run, loading } = useRequest(SendMessages, {
    manual: true,
    onSuccess: () => {
      messageApi.success('提交成功');
      reload?.();
    },
    onError: () => {
      messageApi.error('提交失败，请重试！');
    },
  });
  const [detail, setDetail] = useState('');
  const [detailBody, setDetailBody] = useState('');
  const optionDetail = (data: React.SetStateAction<string>) => {
    setDetail(data)
  };
  return (
    <>
      {contextHolder}
      <ModalForm
       title='发送消息'
        trigger={
          <Button type="primary" icon={<PlusOutlined />}>
            新建
          </Button>
        }
       width="900px"
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (value) => {
          value.status = value.status?1:0;
          value.content=detail
          await run({ data: value as Messages });

          return true;
        }}
      >

        <ProForm.Group>
          <ProFormText
            width="md"
            name="title"
            label="消息主题"
            placeholder="请输入"
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />

        </ProForm.Group>

        <ProForm.Group>
          <WangEditor optionDetail={optionDetail} detailBody={detailBody}/>
        </ProForm.Group>


      </ModalForm>
    </>
  );
};

export default CreateForm;
