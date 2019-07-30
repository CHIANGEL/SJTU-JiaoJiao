import React, { Component } from 'react';
import {
    Alert,
    Text,
    TextInput,
    View,
    StyleSheet,
    Dimensions,
    ScrollView,
    TouchableOpacity,
    Image,
    NativeModules
} from 'react-native';
import {Avatar, Badge, Button, ButtonGroup, SearchBar} from 'react-native-elements';
import Config from '../../Config';
import {NavigationActions} from "react-navigation";
import Textarea from 'react-native-textarea';
import {TimeStamptoDate, TimeStampNow, DatetoTimeStamp} from "../../Utils/TimeStamp";
import Video from "react-native-video";

const {width, height, scale} = Dimensions.get('window');

let ImagePicker = NativeModules.ImageCropPicker;

export default class ReleaseScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedIndex: 0,
            goodName: '',
            goodDiscription: '',
            Price: '',
            images: null,
        };
        this.updateIndex = this.updateIndex.bind(this);
        this.keyID = 0;
    };

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>发布交易信息</Text>)
    };

    updateIndex (selectedIndex) {
        this.setState({
            selectedIndex: selectedIndex,
        })
    }

    updateGoodName = (goodName) => {
        this.setState({
            goodName: goodName,
        })
    };

    updateDiscription = (goodDiscription) => {
        this.setState({
            goodDiscription: goodDiscription,
        })
    };

    updatePrice = (Price) => {
        this.setState({
            Price: Price,
        })
    };

    pickMultiple() {
        ImagePicker.openPicker({
            multiple: true,
            waitAnimationEnd: false,
            includeExif: true,
            forceJpg: true,
        }).then(selectedImages => {
            this.keyID = 0;
            this.setState(previousState => {
                return {
                    images: previousState.images === null ?
                        selectedImages.map(i => {
                            console.log('received image', i);
                            return {uri: i.path, width: i.width, height: i.height, mime: i.mime};
                        }) :
                        previousState.images.concat(selectedImages.map(i => {
                            console.log('received image', i);
                            return {uri: i.path, width: i.width, height: i.height, mime: i.mime};
                        })),
                }});
        }).catch(e => {
            console.warn('出错啦！');
            console.warn(e);
        });
    }

    deleteImage = () => {
        console.warn('删除图片');
    };

    renderVideo(video) {
        console.log('rendering video');
        return (<View style={{height: 300, width: 300}}>
            <Video source={{uri: video.uri, type: video.mime}}
                   style={{position: 'absolute',
                       top: 0,
                       left: 0,
                       bottom: 0,
                       right: 0
                   }}
                   rate={1}
                   paused={false}
                   volume={1}
                   muted={false}
                   resizeMode={'cover'}
                   onError={e => console.log(e)}
                   onLoad={load => console.log(load)}
                   repeat={true} />
        </View>);
    }

    renderImage(image) {
        return (
            <TouchableOpacity>
                <Image style={{
                    width: (1.05 / 5.05 * width),
                    height: (1.05 / 5.05 * width),
                    resizeMode: 'contain',
                    borderColor: 'white',
                    borderWidth: 1,
                }} source={image} />
            </TouchableOpacity>
        )
    }

    renderAsset(image) {
        if (image.mime && image.mime.toLowerCase().indexOf('video/') !== -1) {
            return this.renderVideo(image);
        }

        return this.renderImage(image);
    }

    render() {
        const buttons = ['出售信息', '求购信息'];
        const { selectedIndex } = this.state;

        if (Config.userInfo.userID === -1) {
            Alert.alert(
                '未登录',
                '无法在未登录状态下发布信息，是否切换至登录界面？',
                [
                    {
                        text: '取消',
                        onPress: () => {
                            this.props.navigation.navigate('Home');
                        },
                        style: 'cancel',
                    },
                    {
                        text: '确定', onPress: () => {
                            this.props.navigation.navigate('Login');
                        }
                    },
                ],
                {cancelable: false},
            );
            return (<View/>);
        } else {
            return (
                <ScrollView>
                    <View style={{backgroundColor: '#EFEFF5'}}>
                        <ButtonGroup
                            onPress={this.updateIndex}
                            selectedIndex={selectedIndex}
                            buttons={buttons}
                            containerStyle={{height: 35}}
                            disabledTextStyle={{fontSize: 17}}
                            selectedTextStyle={{fontSize: 17, fontWeight: 'bold'}}
                        />
                        <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                            <Button
                                title='物品名称'
                                titleStyle={{color: 'steelblue', fontSize: 14, fontWeight: 'bold'}}
                                containerStyle={{flex: 1.1}}
                                buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                                raised={true}
                                onPress={() => {alert(this.state.goodName)}}
                            />
                            <SearchBar
                                searchIcon={false}
                                clearIcon={false}
                                cancelIcon={false}
                                lightTheme={true}
                                containerStyle={{
                                    flex: 4,
                                    height: 55,
                                    backgroundColor: 'white',
                                    borderWidth: 0, //no effect
                                    shadowColor: 'white', //no effect
                                    borderBottomColor: 'transparent',
                                    borderTopColor: 'transparent',
                                }}
                                inputContainerStyle={{
                                    marginTop: -9,
                                    height: 55,
                                    backgroundColor: 'white',
                                }}
                                placeholder={'必填。输入物品名称'}
                                onChangeText={ this.updateGoodName }
                                value={ this.state.goodName }
                            />
                        </View>
                        <View style={{height: 10}} />
                        <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                            <Button
                                title='详细描述'
                                titleStyle={{color: 'steelblue', fontSize: 14, fontWeight: 'bold'}}
                                containerStyle={{flex: 1.05}}
                                buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                                raised={true}
                                onPress={() => {alert(this.state.goodDiscription)}}
                            />
                            <Textarea
                                containerStyle={{
                                    flex: 4,
                                    height: 150,
                                    backgroundColor: 'white',
                                    borderWidth: 0, //no effect
                                    shadowColor: 'white', //no effect
                                    borderBottomColor: 'transparent',
                                    borderTopColor: 'transparent',
                                }}
                                style={{marginLeft: 14, marginRight: 13, fontSize: 18, }}
                                onChangeText={this.updateDiscription}
                                defaultValue={this.state.goodDiscription}
                                maxLength={360}
                                placeholder={'选填。可输入对物品的详细描述或者列出物品清单'}
                                placeholderTextColor={'grey'}
                                underlineColorAndroid={'transparent'}
                            />
                        </View>
                        <View style={{height: 10}} />
                        <View>
                            <Text style={{alignItems: 'center', justifyContent: 'center'}}>
                                商品标签
                            </Text>
                        </View>
                        <View style={{height: 10}} />
                        <View style={{flexDirection: 'row'}}>
                            <TouchableOpacity onPress={this.pickMultiple.bind(this)}>
                                <Image style={{width: (1.05 / 5.05 * width), height: (1.05 / 5.05 * width), resizeMode: 'contain',}} source={require('../../assets/images/addPhotoVideo.jpg')}/>
                            </TouchableOpacity>
                            <ScrollView horizontal={true}>
                                {this.state.images ? this.state.images.map(i => <View key={this.keyID++}>{this.renderAsset(i)}</View>) : null}
                            </ScrollView>
                        </View>
                        <View style={{height: 10}} />
                        <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                            <Button
                                title='预期价格'
                                titleStyle={{color: 'steelblue', fontSize: 14, fontWeight: 'bold'}}
                                containerStyle={{flex: 1.1}}
                                buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                                raised={true}
                                onPress={() => {alert(this.state.Price)}}
                            />
                            <SearchBar
                                searchIcon={false}
                                clearIcon={false}
                                cancelIcon={false}
                                lightTheme={true}
                                containerStyle={{
                                    flex: 4,
                                    height: 55,
                                    backgroundColor: 'white',
                                    borderWidth: 0, //no effect
                                    shadowColor: 'white', //no effect
                                    borderBottomColor: 'transparent',
                                    borderTopColor: 'transparent',
                                }}
                                inputContainerStyle={{
                                    marginTop: -9,
                                    height: 55,
                                    backgroundColor: 'white',
                                }}
                                placeholder={'选填'}
                                onChangeText={ this.updatePrice }
                                value={ this.state.Price }
                            />
                        </View>
                        <View style={{height: 10}} />
                        <Button
                            title={'发布信息'}
                            titleStyle={{color: 'white', fontSize: 18}}
                            buttonStyle={{backgroundColor: 'red'}}
                            containerStyle={{width: 160, marginLeft: (width / 2 - 80)}}
                            raised={true}
                            onPress={() => Alert.alert(
                                '发布信息',
                                ('您确定要发布该条' + buttons[this.state.selectedIndex] + '吗？'),
                                [
                                    {
                                        text: '取消',
                                        onPress: () => console.warn('Cancel Pressed'),
                                        style: 'cancel',
                                    },
                                    {
                                        text: '确定', onPress: () => {
                                            let addType;
                                            if (this.state.selectedIndex === 0)
                                                addType = 'sellInfo';
                                            else
                                                addType = 'buyInfo';
                                            console.warn((Config.fetchPrefix + addType));
                                            fetch((Config.fetchPrefix + addType), {
                                                method: 'POST',
                                                headers: {
                                                    Accept: 'application/json',
                                                    'Content-Type': 'application/x-www-form-urlencoded',
                                                    Authorization: ('Bearer ' + Config.JaccountToken.token),
                                                },
                                                body: ('userID=' + Config.userInfo.userID + '&validTime=20&goodName=' + this.state.goodName),
                                            })
                                                .then((response) => {
                                                    if(response.ok) {
                                                        Alert.alert(
                                                            '发布成功',
                                                            '成功发布该交易信息：' + this.state.goodName,
                                                            [
                                                                {text: '好', onPress: () => {}}
                                                            ],
                                                            {cancelable: false},
                                                        )
                                                        this.setState({
                                                            goodName: '',
                                                            goodDiscription: '',
                                                            Price: '',
                                                        });
                                                    } else {
                                                        Alert.alert(
                                                            '出错啦',
                                                            '网络可能出了问题，请再试一次吧',
                                                            [
                                                                {text: '好', onPress: () => {}}
                                                            ],
                                                            {cancelable: false},
                                                        )
                                                    }
                                                })
                                                .catch((error) => {
                                                    console.error(error);
                                                });
                                        }
                                    },
                                ],
                                {cancelable: false},
                            )}
                        />
                        <View style={{height: 10}} />
                    </View>
                </ScrollView>
            )
        }
    }
}
