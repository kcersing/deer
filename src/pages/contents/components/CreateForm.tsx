import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm,
  ProForm,
  ProFormDateRangePicker,ProFormSelect,ProFormDigit,
  ProFormText,
  ProFormTextArea,
  ProFormUploadButton,ProFormList,ProCard,ProFormDependency,
  ProFormSwitch,
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message,Descriptions, } from 'antd';
import React, { FC,useState } from 'react';

import { createArticle }from "@/pages/contents/service/service";
import { Article } from  "@/pages/contents/service/data";
import dayjs from 'dayjs';
import type { Dayjs } from 'dayjs';
import WangEditor from '@/pages/components/wangeditor'
import {getDicthtList} from "../../system/dict/service/service";
interface CreateFormProps {
  reload?: ActionType['reload'];
}

const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const { data: tagEnum } = useRequest(() => {
    return  getDicthtList({'dictId':2})
  });

  const [messageApi, contextHolder] = message.useMessage();

  const defaultValue = [dayjs().startOf('month'), dayjs().startOf('month').add(2, 'month').subtract(1, 'day'),];

  const { run, loading } = useRequest(createArticle, {
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
       title='新建'
        trigger={
          <Button type="primary" icon={<PlusOutlined />}>新建</Button>
        }
        width="800px"
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (values) => {
          values.status = values.status?1:0;
          // value.content= detail;
          await run({ data: values as Article });
          return true;
        }}
      >

        <ProForm.Group>
          <ProFormUploadButton
            name="pic"
            label="图片"
            max={1}
            fieldProps={{
              name: 'file',
              listType: 'picture-card',
            }}
            action="/upload.do"
            extra="图片不能大于1M"
          />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText
            width="md"
            name="name"
            label="名称"
            tooltip="最长为 24 位"
            placeholder="请输入名称"
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormSelect
            width="md"
            name="tagId"
            label="标签"
            options= {tagEnum}
            fieldProps= {
              {
                showSearch: true,
                multiple: true,
                mode:"multiple",
                fieldNames:{ label: 'title', value: 'id' }
              }
            }
          />
        </ProForm.Group>

        <ProForm.Group>
            <ProFormSwitch
              name="status"
              width="md"
              label="状态"
              checkedChildren="发布"
              unCheckedChildren="下架"
              initialValue="发布"
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
