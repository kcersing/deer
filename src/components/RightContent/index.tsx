import { QuestionCircleOutlined } from '@ant-design/icons';
import { SelectLang as UmiSelectLang } from '@umijs/max';
import { useLocation } from 'umi';
export type SiderTheme = 'light' | 'dark';

export const SelectLang: React.FC = () => {
  return ( <></>
    // <UmiSelectLang
    //   style={{
    //     padding: 4,
    //   }}
    // />
  );
};

export const Question: React.FC = () => {

  return (
    <a
      href="http://www.kkksy.com"
      target="_blank"
      rel="noreferrer"
      style={{
        display: 'inline-flex',
        padding: '4px',
        fontSize: '18px',
        color: 'inherit',
      }}
    >
      <QuestionCircleOutlined />
    </a>
  );
};
