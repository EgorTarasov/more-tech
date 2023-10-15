import { Col, Rate, Row } from 'antd';
import { IAtm } from '../api/models/IAtm';
import distanceConverter from '../utils/distanceConverter';

type Props = {
    atm: IAtm;
};

const AtmGeneral = ({ atm }: Props) => {
    return (
        <>
            <div className='department-general'>
                <Row>
                    <div className='department-general__name department-general__item'>
                        Банкомат
                    </div>
                </Row>

                <Row wrap={false} justify={'space-between'}>
                    <Col span={18}>
                        <div className='department-general__address department-general__item'>
                            {atm.address}
                        </div>
                    </Col>

                    <Col>
                        <div className='department-general__distance department-general__item'>
                            {distanceConverter(atm.distance)}
                        </div>
                    </Col>
                </Row>

                <Row justify={'space-between'} align={'middle'} style={{ marginTop: 7 }}>
                    <Col>
                        <Rate allowHalf defaultValue={5} />
                    </Col>
                </Row>
            </div>
        </>
    );
};

export default AtmGeneral;
