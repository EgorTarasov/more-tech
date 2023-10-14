import { Button, Col, Rate, Row } from 'antd';
import { IDepartment } from '../api/models';
import distanceConverter from '../utils/distanceConverter';
import { FileAddOutlined } from '@ant-design/icons';
import { useStores } from '../hooks/useStores';

type Props = {
    department: IDepartment;
};

const DepartmentGeneral = ({ department }: Props) => {
    const { rootStore } = useStores();

    return (
        <div
            onClick={() => {
                rootStore.setSelectedDepartment(department);
            }}
            className='department-general'
        >
            <Row>
                <div className='department-general__name department-general__item'>
                    {department.shortName}
                </div>
            </Row>

            <Row>
                <div className='department-general__wait-time department-general__item'>
                    Ожидание - 5 минут
                </div>
            </Row>

            <Row wrap={false} justify={'space-between'}>
                <Col span={18}>
                    <div className='department-general__address department-general__item'>
                        {department.address}
                    </div>
                </Col>

                <Col>
                    <div className='department-general__distance department-general__item'>
                        {distanceConverter(department.distance)}
                    </div>
                </Col>
            </Row>

            <Row justify={'space-between'} align={'middle'} style={{ marginTop: 7 }}>
                <Col>
                    <Rate allowHalf defaultValue={4.5} />
                </Col>

                <Col>
                    В избранное
                    <Button
                        style={{
                            marginLeft: 10,
                            backgroundColor: '#EBEDF5',
                            color: '#6b7683',
                            border: 'none',
                        }}
                        type='default'
                        shape='circle'
                        icon={<FileAddOutlined />}
                    />
                </Col>
            </Row>
        </div>
    );
};

export default DepartmentGeneral;
