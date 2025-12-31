import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm,
  ProForm,
  ProFormDateRangePicker,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
  ProFormUploadButton,
  ProFormList,
  ProCard,
  ProFormDependency,
  ProFormSwitch,
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message,Descriptions, } from 'antd';

import React, {FC,useEffect, useRef, useState} from 'react';
import { createProduct }from "@/pages/product/list/service/service";
import { Product } from  "@/pages/product/list/service/data";
import dayjs from 'dayjs';
import type { Dayjs } from 'dayjs';
import {getItemList,getItem} from "@/pages/product/item/service/service";

import CreateForm from "./CreateForm";
import {queryFakeList} from "../../../account/center/service";

interface CreateFormProps {
  reload?: ActionType['reload'];
}

const AttributesItme = (props) => {
  const { id } = props;
  if (!(id >= 0)) {
    return null;
  }
  const [itemData, setItemData] = useState();

  useEffect(() => {
    const getResponses = async () => {
      const [res] = await Promise.all([
        getItem({id:id}),
      ]);
      setItemData(res.data)
    }
    getResponses();
  }, [id]);


  if (!itemData){
    return null;
  }

  return (
      <Descriptions  size="small" column={2}>
        <Descriptions.Item label="名称">{itemData.name}</Descriptions.Item>
        <Descriptions.Item label="Code">{itemData.code}</Descriptions.Item>
        <Descriptions.Item label="类别">{itemData.type}</Descriptions.Item>
        <Descriptions.Item label="金额">{itemData.price}</Descriptions.Item>

        <Descriptions.Item label="期限">{itemData.duration}</Descriptions.Item>
        <Descriptions.Item label="次数">{itemData.count}</Descriptions.Item>
        <Descriptions.Item label="时长">{itemData.length}</Descriptions.Item>
        <Descriptions.Item label="标签">{itemData.tagName}</Descriptions.Item>

        <Descriptions.Item label="备注">
        {itemData.desc}
        </Descriptions.Item>
      </Descriptions>
  );
}

const Attributes  = (props) => {
  return (
    <ProFormList
      name="attributes"
      label="属性"
      creatorButtonProps={{
        creatorButtonText: '添加属性项',
      }}
      min={1}
      copyIconProps={false}
      itemRender={({ listDom, action }, { index }) => (
        <ProCard
          variant="outlined"
          style={{ marginBlockEnd: 8 }}
          title={`属性${index + 1}`}
          extra={action}
          styles={{ body: { paddingBlockEnd: 0 }} }
          bodyStyle={{padding: "unset"}}
        >
          {listDom}
        </ProCard>
      )}
    >

      <ProCard

        split="vertical"


      >
        <ProCard bodyStyle={{ width: 240,flexShrink: 0 }}>
          <ProFormSelect
            name="itme"
            label="选择属性"
            width="md"
            params={{current: 999, pageSize: 1}}
            request={(params)=>{
              return getItemList({params}).then((res) => {return res.data})
            }}

            fieldProps={{
              fieldNames: {
                label: 'name',
                value: 'id',
              },
            }}
            placeholder="请选择"
            rules={[{ required: true, message: '请选择!' }]}
          />
        </ProCard>

        <ProCard  bodyStyle={{ width: 500,flexShrink: 0,padding: "unset",paddingLeft: 24 }}>
          <ProFormDependency name={["itme"]} ignoreFormListField={false} >
            {(depValues) => (

              <AttributesItme id={depValues["itme"]} />
            )}
          </ProFormDependency>
        </ProCard>
      </ProCard>
    </ProFormList>


  );

}
export default Attributes;
